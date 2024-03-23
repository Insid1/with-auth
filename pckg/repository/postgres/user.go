package postgres

import "user"

// todo: repo is not implemented. Pls connect real database and create queries.

type Postgres struct{}

func NewUserPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Create(u *user.User) (uint64, error) {
	return 1, nil
}

func (p *Postgres) Get(id uint64) (*user.User, error) {
	return &user.User{
		Id:       0,
		Name:     "my name",
		Email:    "my email",
		Password: "hashed password",
	}, nil
}

func (p *Postgres) Update(u *user.User) (*user.User, error) {
	return &user.User{
		Id:       0,
		Name:     "my updated name",
		Email:    "my updated email",
		Password: "updated hashed password",
	}, nil
}
func (p *Postgres) Delete(id uint64) error {
	return nil
}
