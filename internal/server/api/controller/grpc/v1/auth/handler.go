package v1

import (
	pb "github.com/ivas1ly/gophkeeper/pkg/api/v1/auth"
	"go.uber.org/zap"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	log *zap.Logger
}

var _ pb.AuthServiceServer = (*AuthHandler)(nil)

func NewAuthHandler(log *zap.Logger) *AuthHandler {
	return &AuthHandler{
		log: log,
	}
}
