package user

import (
	"database/sql"

	"github.com/Insid1/go-auth-user/user-service/internal/model"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Get(id string) (*model.User, error) {
	var usr model.User

	rows, err := r.DB.Query("SELECT * FROM \"user\" WHERE id=$1", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if rows.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Age) != nil {
			return nil, err
		}
	}

	return &usr, nil
}

func (r *Repository) Create(usr *model.User) (string, error) {
	var id string
	err := r.DB.QueryRow(
		"INSERT INTO \"user\" (name, email, age) VALUES ($1, $2, $3) RETURNING id;",
		usr.Name, usr.Email, usr.Age).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
