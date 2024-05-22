package repository

import (
	"database/sql"

	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/repository/user"
)

type Repository struct {
	UserRepository
}

type UserRepository interface {
	Get(string) (*model.User, error)
	Create(*model.User) (string, error)
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: user.NewRepository(db),
	}
}
