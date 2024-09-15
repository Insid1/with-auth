package service

import (
	"github.com/Insid1/go-auth-user/user/internal/model"
	"github.com/Insid1/go-auth-user/user/internal/repository"
	"github.com/Insid1/go-auth-user/user/internal/service/user"
)

type User interface {
	Get(id string, email string) (*model.User, error)
	Create(*model.User) (string, error)
}

func NewUserService(repo repository.User) User {
	return &user.Service{
		UserRepository: repo,
	}
}
