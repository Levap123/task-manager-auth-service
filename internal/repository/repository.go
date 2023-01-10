package repository

import (
	"context"

	"github.com/Levap123/task-manager-auth-service/internal/repository/postgres"
	"github.com/Levap123/task-manager-auth-service/proto"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth
}

type Auth interface {
	Create(ctx context.Context, user *proto.User) (*proto.User, error)
	Get(ctx context.Context, email string) (*proto.User, error)
}

func NewRepoPostgres(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: postgres.NewAuthRepo(db),
	}
}
