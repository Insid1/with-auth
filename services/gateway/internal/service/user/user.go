package user

import (
	"github.com/Insid1/go-auth-user/gateway/internal/entity"
	"github.com/Insid1/go-auth-user/gateway/internal/repository"
)

type Service struct {
	Repo repository.User
}

func (p *Service) Create(u *entity.User) (string, error) {
	return p.Repo.Create(u)
}

func (p *Service) Get(id string) (*entity.User, error) {
	return p.Repo.Get(id)
}

func (p *Service) Update(u *entity.User) (*entity.User, error) {
	return p.Repo.Update(u)
}

func (p *Service) Delete(id string) error {
	return p.Repo.Delete(id)
}
