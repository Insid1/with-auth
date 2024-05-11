package postgres

import (
	"goAuth/internal/entity"
)

type Postgres struct{}

func NewUserPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Create(u *entity.User) (uint64, error) {
	return 1, nil
}

func (p *Postgres) Get(id uint64) (*entity.User, error) {
	return &entity.User{
		Id:       0,
		Name:     "my name",
		Email:    "my email",
		Password: "hashed password",
	}, nil
}

func (p *Postgres) Update(u *entity.User) (*entity.User, error) {
	return &entity.User{
		Id:       0,
		Name:     "my updated name",
		Email:    "my updated email",
		Password: "updated hashed password",
	}, nil
}
func (p *Postgres) Delete(id uint64) error {
	return nil
}
