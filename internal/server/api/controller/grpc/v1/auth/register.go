package v1

import (
	"context"

	pb "github.com/ivas1ly/gophkeeper/pkg/api/v1/auth"
)

func (ah *AuthHandler) Register(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	ah.log.Info("Register")
	return nil, nil
}
