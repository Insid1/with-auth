package service

import "user"

type User interface {
	Create(*user.User) (uint64, error)
	Get(id uint64) (*user.User, error)
	Update(id uint64) (*user.User, error)
	Delete(id uint64) error
}
