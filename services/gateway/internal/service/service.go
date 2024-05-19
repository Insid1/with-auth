package service

import (
	"github.com/Insid1/go-auth-user/gateway/internal/entity"
	"github.com/Insid1/go-auth-user/gateway/internal/repository"
	"github.com/Insid1/go-auth-user/gateway/internal/service/user"
)

type User interface {
	Create(*entity.User) (string, error)
	Get(string) (*entity.User, error)
	Update(*entity.User) (*entity.User, error)
	Delete(string) error
}

func NewUserService() User {
	repo := repository.NewRepository()

	return &user.Service{Repo: repo}
}
