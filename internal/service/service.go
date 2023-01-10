package service

import (
	"context"

	"github.com/Levap123/task-manager-auth-service/internal/repository"
	"github.com/Levap123/task-manager-auth-service/proto"
)

type Service struct {
	Auth
}

type Auth interface {
	SignIn(ctx context.Context, name, password string) (string, error)
	Create(ctx context.Context, user *proto.User) (*proto.User, error)
}


func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.Auth),
	}
}

