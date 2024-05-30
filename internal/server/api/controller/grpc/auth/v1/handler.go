package v1

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/ivas1ly/gophkeeper/internal/server/entity"
	authpb "github.com/ivas1ly/gophkeeper/pkg/api/gophkeeper/auth/v1"
	"go.uber.org/zap"
)

type AuthService interface {
	Register(ctx context.Context, username, password string) (*entity.User, error)
	Login(ctx context.Context, username, password string) (*entity.User, error)
}

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	log         *zap.Logger
	authService AuthService
	validator   *protovalidate.Validator
}

var _ authpb.AuthServiceServer = (*AuthHandler)(nil)

func NewAuthHandler(authService AuthService, validator *protovalidate.Validator) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		log:         zap.L().With(zap.String("handler", "auth")),
		validator:   validator,
	}
}
