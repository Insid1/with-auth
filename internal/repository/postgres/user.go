package postgres

import (
	"github.com/Insid1/go-auth-user/internal/entity"
)

type Postgres struct{}

func NewUserPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Create(u *entity.User) (string, error) {
	return u.Id, nil
}

func (p *Postgres) Get(id string) (*entity.User, error) {
	return &entity.User{
		Id:       "123",
		Name:     "my name",
		Email:    "my email",
		Password: "hashed password",
	}, nil
}

func (p *Postgres) Update(u *entity.User) (*entity.User, error) {
	return &entity.User{
		Id:       "123",
		Name:     "my updated name",
		Email:    "my updated email",
		Password: "updated hashed password",
	}, nil
}
func (p *Postgres) Delete(id string) error {
	return nil
}
