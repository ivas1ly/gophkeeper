package server

import (
	"context"
	"log"
	"net"

	"github.com/ivas1ly/gophkeeper/internal/lib/logger"
	authV1 "github.com/ivas1ly/gophkeeper/internal/server/api/controller/grpc/v1/auth"
	"github.com/ivas1ly/gophkeeper/internal/server/config"
	authDesc "github.com/ivas1ly/gophkeeper/pkg/api/v1/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"go.uber.org/zap"
)

type Server struct {
	grpcServer *grpc.Server
	log        *zap.Logger
	cfg        config.Config
}

func New(_ context.Context, cfg config.Config) (*Server, error) {
	log := logger.New(cfg.Server.LogLevel, logger.NewDefaultLoggerConfig()).
		With(zap.String("server", "gophkeeper"))
	logger.SetGlobalLogger(log)

	a := &Server{
		log: log,
		cfg: cfg,
	}

	return a, nil
}

func (s *Server) Run(ctx context.Context) error {
	defer func() {
		err := s.log.Sync()
		if err != nil {
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
	s.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(s.grpcServer)

	authHandler := authV1.NewAuthHandler(s.log)

	authDesc.RegisterAuthServiceServer(s.grpcServer, authHandler)
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
