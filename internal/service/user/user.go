package service

import (
	"goAuth/internal/entity"
	"goAuth/internal/repository"
	"goAuth/internal/service"
)

type UserService struct {
	repo repository.User
}

func NewUserService() service.User {
	repo := repository.NewRepository()
	return &UserService{repo: repo}
}

func (p *UserService) Create(u *entity.User) (uint64, error) {
	return p.repo.Create(u)
}

func (p *UserService) Get(id uint64) (*entity.User, error) {
	return p.repo.Get(id)
}

func (p *UserService) Update(u *entity.User) (*entity.User, error) {
	return p.repo.Update(u)
}

func (p *UserService) Delete(id uint64) error {
	return p.repo.Delete(id)
}
