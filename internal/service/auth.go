package service

import (
	"context"

	"github.com/Levap123/task-manager-auth-service/internal/repository"
	pass "github.com/Levap123/task-manager-auth-service/pkg/password"
	"github.com/Levap123/task-manager-auth-service/proto"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (as *AuthService) Create(ctx context.Context, user *proto.User) (*proto.User, error) {
	user.Password = pass.GeneratePasswordHash(user.Password)
	return as.repo.Create(ctx, user)
}
