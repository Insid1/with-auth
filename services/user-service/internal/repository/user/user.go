package user

import (
	"database/sql"
	"fmt"

	"github.com/Insid1/go-auth-user/user-service/internal/model"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Get(id string) (*model.User, error) {
	var usr model.User

	err := r.DB.QueryRow(
		"SELECT id, name, email, age, created_at, updated_at FROM \"user\" WHERE id=$1", id).Scan(
		&usr.ID, &usr.Name, &usr.Email, &usr.Age, &usr.CreatedAt, &usr.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *Repository) GetBy(column string, source string) (*model.User, error) {
	var usr model.User
	query := fmt.Sprintf("SELECT id, name, email, age, password,created_at, updated_at FROM \"user\" WHERE %s=$1", column)

	err := r.DB.QueryRow(query, source).Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Age, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *Repository) Create(usr *model.User) (string, error) {
	var id string
	err := r.DB.QueryRow(
		"INSERT INTO \"user\" (name, email, age, password) VALUES ($1, $2, $3, $4) RETURNING id;",
		usr.Name, usr.Email, usr.Age, usr.Password).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
