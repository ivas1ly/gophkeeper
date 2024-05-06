package v1

import (
	authpb "github.com/ivas1ly/gophkeeper/pkg/api/gophkeeper/auth/v1"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	log *zap.Logger
}

var _ authpb.AuthServiceServer = (*AuthHandler)(nil)

func NewAuthHandler(log *zap.Logger) *AuthHandler {
	return &AuthHandler{
		log: log,
	}
}
