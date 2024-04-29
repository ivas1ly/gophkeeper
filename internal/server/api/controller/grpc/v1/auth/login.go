package v1

import (
	"context"

	pb "github.com/ivas1ly/gophkeeper/pkg/api/v1/auth"
)

func (ah *AuthHandler) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	ah.log.Info("Login")
	return nil, nil
}
