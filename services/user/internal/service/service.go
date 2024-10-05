package service

import (
	"github.com/Insid1/with-auth/user/internal/model"
	"github.com/Insid1/with-auth/user/internal/repository"
	"github.com/Insid1/with-auth/user/internal/service/user"
)

type User interface {
	Get(id string, email string) (*model.User, error)
	Create(usr *model.User, password string) (*model.User, error)
	Update(usr *model.User, password string) (*model.User, error)
	CheckPassword(id string, email, password string) (*model.User, error)
	Delete(id string) (string, error)
}

func NewUserService(repo repository.User) User {
	return &user.Service{
		UserRepository: repo,
	}
}
