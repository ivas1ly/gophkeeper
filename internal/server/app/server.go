package server

import (
	"context"
	"errors"
	"log"
	"net"
	"syscall"

	"github.com/bufbuild/protovalidate-go"
	"github.com/ivas1ly/gophkeeper/internal/lib/logger"
	"github.com/ivas1ly/gophkeeper/internal/lib/storage/postgres"
	authV1 "github.com/ivas1ly/gophkeeper/internal/server/api/controller/grpc/auth/v1"
	reqlogger "github.com/ivas1ly/gophkeeper/internal/server/api/interceptor"
	"github.com/ivas1ly/gophkeeper/internal/server/app/provider"
	"github.com/ivas1ly/gophkeeper/internal/server/config"
	authV1Desc "github.com/ivas1ly/gophkeeper/pkg/api/gophkeeper/auth/v1"
	"github.com/ivas1ly/gophkeeper/pkg/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"go.uber.org/zap"
)

type Server struct {
	grpcServer      *grpc.Server
	log             *zap.Logger
	db              *postgres.DB
	serviceProvider *provider.ServiceProvider
	validator       *protovalidate.Validator
	cfg             config.Config
}

func New(ctx context.Context, cfg config.Config) (*Server, error) {
	log := logger.New(cfg.App.LogLevel, logger.NewDefaultLoggerConfig()).
		With(zap.String("server", "gophkeeper"))
	logger.SetGlobalLogger(log)

	s := &Server{
		log: log,
		cfg: cfg,
	}

	jwt.SigningKey = cfg.App.SigningKey
	jwt.ExpirationTime = cfg.App.ExpirationTime
	jwt.Issuer = cfg.App.Name

	s.log.Info("init the database pool")
	db, err := postgres.New(ctx, cfg.DatabaseURI, cfg.DatabaseConnAttempts, cfg.DatabaseConnTimeout)
	if err != nil {
		s.log.Error("can't establish a connection to the database", zap.Error(err))
		return nil, err
	}

	s.log.Info("database connection established")
	s.db = db

	s.log.Info("init proto validator")
	s.validator, err = protovalidate.New()
	if err != nil {
		s.log.Error("failed to initialize validator", zap.Error(err))
		return nil, err
	}

	s.log.Info("init services")
	s.serviceProvider = provider.NewServiceProvider(db)
	s.serviceProvider.RegisterServices()

	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	defer func() {
		err := s.log.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) {
			log.Println(err)
		}
	}()

	s.initGRPCServer(ctx)
	err := s.runGRPCServer()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) initGRPCServer(_ context.Context) {
	s.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(reqlogger.NewInterceptor(s.log)),
	)

	reflection.Register(s.grpcServer)

	authV1Handler := authV1.NewAuthHandler(s.serviceProvider.AuthService, s.validator)

	authV1Desc.RegisterAuthServiceServer(s.grpcServer, authV1Handler)
}

func (s *Server) runGRPCServer() error {
	listen, err := net.Listen("tcp", s.cfg.RunAddress)
	if err != nil {
		return err
	}

	s.log.Info("gRPC server started", zap.String("addr", s.cfg.RunAddress))
	err = s.grpcServer.Serve(listen)
	if err != nil {
		s.log.Error("gRPC server", zap.Error(err))
		return err
	}

	return nil
}
