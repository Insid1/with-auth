package repository

import (
	"context"
	"database/sql"

	"github.com/Insid1/go-auth-user/user/internal/model"
	"github.com/Insid1/go-auth-user/user/internal/repository/user"
)

type User interface {
	Get(string) (*model.User, error)
	GetBy(string, string) (*model.User, error)
	Create(*model.User) (string, error)
}

func NewUserRepository(ctx context.Context, db *sql.DB) User {
	return &user.Repository{DB: db}
}
