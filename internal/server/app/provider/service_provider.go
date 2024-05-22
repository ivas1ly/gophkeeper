package provider

import (
	"context"

	"github.com/ivas1ly/gophkeeper/internal/lib/storage/postgres"
	"github.com/ivas1ly/gophkeeper/internal/server/entity"
	"github.com/ivas1ly/gophkeeper/internal/server/repository"
	"github.com/ivas1ly/gophkeeper/internal/server/service"
)

type AuthService interface {
	Register(ctx context.Context, username, password string) (*entity.User, error)
	Login(ctx context.Context, username, password string) (*entity.User, error)
}

type AuthRepository interface {
	AddUser(ctx context.Context, userInfo *entity.UserInfo) (*entity.User, error)
	FindUser(ctx context.Context, username string) (*entity.User, error)
}

type ServiceProvider struct {
	AuthService AuthService

	db *postgres.DB
}

func NewServiceProvider(db *postgres.DB) *ServiceProvider {
	return &ServiceProvider{
		db: db,
	}
}

func (s *ServiceProvider) RegisterServices() {
	s.NewAuthService()
}

func (s *ServiceProvider) newAuthRepository() AuthRepository {
	return repository.NewAuthRepository(s.db)
}

func (s *ServiceProvider) NewAuthService() AuthService {
	if s.AuthService == nil {
		s.AuthService = service.NewAuthService(s.newAuthRepository())
	}

	return s.AuthService
}
