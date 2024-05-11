package repository

import (
	"goAuth/internal/entity"
	"goAuth/internal/repository/postgres"
)

type User interface {
	Create(u *entity.User) (uint64, error)
	Get(id uint64) (*entity.User, error)
	Update(u *entity.User) (*entity.User, error)
	Delete(id uint64) error
}

func NewRepository() User {
	return postgres.NewUserPostgres()
}
