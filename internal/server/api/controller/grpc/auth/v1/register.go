package v1

import (
	"context"

	authpb "github.com/ivas1ly/gophkeeper/pkg/api/gophkeeper/auth/v1"
)

func (ah *AuthHandler) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	ah.log.Info("Register")
	return nil, nil
}
