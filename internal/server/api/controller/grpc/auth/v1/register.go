package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/ivas1ly/gophkeeper/internal/server/api/controller"
	"github.com/ivas1ly/gophkeeper/internal/server/entity"
	authpb "github.com/ivas1ly/gophkeeper/pkg/api/gophkeeper/auth/v1"
	"github.com/ivas1ly/gophkeeper/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (ah *AuthHandler) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	if err := ah.validator.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, controller.MsgFromValidationError(err))
	}

	user, err := ah.authService.Register(ctx, req.Username, req.Password)
	if errors.Is(err, entity.ErrUsernameUniqueViolation) {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("username %q already exists", req.Username))
	}
	if err != nil {
		return nil, status.Error(codes.Internal, controller.MsgInternalServerError)
	}

	jwtToken, expiredAt, err := jwt.NewToken(user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, controller.MsgInternalServerError)
	}

	return &authpb.RegisterResponse{
		Id:          user.ID,
		Username:    user.Username,
		AccessToken: jwtToken,
		ExpiredAt:   timestamppb.New(expiredAt),
	}, nil
}
