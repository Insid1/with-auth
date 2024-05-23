package repository

import (
	"context"
	"database/sql"

	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/repository/user"
)

type User interface {
	Get(string) (*model.User, error)
	Create(*model.User) (string, error)
}

func NewUserRepository(ctx context.Context, db *sql.DB) User {
	return &user.Repository{DB: db}
}
