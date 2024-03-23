package repository

import (
	"user"
	"user/pckg/repository/postgres"
)

type User interface {
	Create(u *user.User) (uint64, error)
	Get(id uint64) (*user.User, error)
	Update(u *user.User) (*user.User, error)
	Delete(id uint64) error
}

func NewRepository() User {
	return postgres.NewUserPostgres()
}
