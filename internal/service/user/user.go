package service

import (
	"github.com/Insid1/go-auth-user/internal/entity"
	"github.com/Insid1/go-auth-user/internal/repository"
	"github.com/Insid1/go-auth-user/internal/service"
)

type UserService struct {
	repo repository.User
}

func NewUserService() service.User {
	repo := repository.NewRepository()
	return &UserService{repo: repo}
}

func (p *UserService) Create(u *entity.User) (string, error) {
	return p.repo.Create(u)
}

func (p *UserService) Get(id string) (*entity.User, error) {
	return p.repo.Get(id)
}

func (p *UserService) Update(u *entity.User) (*entity.User, error) {
	return p.repo.Update(u)
}

func (p *UserService) Delete(id string) error {
	return p.repo.Delete(id)
}
