package repository

import (
	"goAuth/internal/entity"
	"goAuth/internal/repository/postgres"
)

type User interface {
	Create(u *entity.User) (string, error)
	Get(id string) (*entity.User, error)
	Update(u *entity.User) (*entity.User, error)
	Delete(id string) error
}

func NewRepository() User {
	return postgres.NewUserPostgres()
}
