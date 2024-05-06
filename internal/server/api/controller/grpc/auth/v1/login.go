package v1

import (
	"context"

	authpb "github.com/ivas1ly/gophkeeper/pkg/api/gophkeeper/auth/v1"
)

func (ah *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	ah.log.Info("Login")
	return nil, nil
}
