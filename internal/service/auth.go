package service

import (
	"context"
	"errors"

	"github.com/Levap123/task-manager-auth-service/internal/repository"
	"github.com/Levap123/task-manager-auth-service/pkg/jwt"
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

func (as *AuthService) SignIn(ctx context.Context, name, password string) (string, error) {
	user, err := as.repo.Get(ctx, name)
	if err != nil {
		return "", err
	}
	if err := pass.ComparePassword(user.Password, password); err != nil {
		return "", errors.New("invalid passsword")
	}
	token, err := jwt.GenerateJwt(int(user.Id))
	if err != nil {
		return "", err
	}
	return token, nil
}
