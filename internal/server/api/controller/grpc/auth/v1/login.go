package v1

import (
	"context"
	"errors"

	"github.com/ivas1ly/gophkeeper/internal/server/api/controller"
	"github.com/ivas1ly/gophkeeper/internal/server/entity"
	authpb "github.com/ivas1ly/gophkeeper/pkg/api/gophkeeper/auth/v1"
	"github.com/ivas1ly/gophkeeper/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (ah *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	if err := ah.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, controller.MsgFromValidationError(err))
	}

	user, err := ah.authService.Login(ctx, req.Username, req.Password)
	if errors.Is(err, entity.ErrUsernameNotFound) {
		return nil, status.Error(codes.Unauthenticated, entity.ErrIncorrectLoginOrPassword.Error())
	}
	if errors.Is(err, entity.ErrIncorrectLoginOrPassword) {
		return nil, status.Error(codes.Unauthenticated, entity.ErrIncorrectLoginOrPassword.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Internal, controller.MsgInternalServerError)
	}

	jwtToken, expiredAt, err := jwt.NewToken(user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, controller.MsgInternalServerError)
	}

	return &authpb.LoginResponse{
		AccessToken: jwtToken,
		ExpiredAt:   timestamppb.New(expiredAt),
	}, nil
}
