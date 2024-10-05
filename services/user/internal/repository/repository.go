package repository

import (
	"database/sql"

	"github.com/Insid1/with-auth/user/internal/model"
	"github.com/Insid1/with-auth/user/internal/repository/user"
)

type User interface {
	Get(id string) (*model.User, error)
	GetBy(column string, source string) (*model.User, error)
	Create(usr *model.User) (*model.User, error)
	Update(usr *model.User) (*model.User, error)
	Delete(id string) (string, error)
}

func NewUserRepository(db *sql.DB) User {
	return &user.Repository{DB: db}
}
