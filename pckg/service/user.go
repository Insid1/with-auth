package service

import (
	"user"
	"user/pckg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewRepository(),
	}
}

func (p *UserService) Create(u *user.User) (uint64, error) {
	return p.repo.Create(u)
}

func (p *UserService) Get(id uint64) (*user.User, error) {
	return p.repo.Get(id)
}

func (p *UserService) Update(u *user.User) (*user.User, error) {
	return p.repo.Update(u)
}

func (p *UserService) Delete(id uint64) error {
	return p.repo.Delete(id)
}
