package service

import (
	"goAuth/internal/entity"
)

type User interface {
	Create(*entity.User) (uint64, error)
	Get(uint64) (*entity.User, error)
	Update(*entity.User) (*entity.User, error)
	Delete(uint64) error
}
