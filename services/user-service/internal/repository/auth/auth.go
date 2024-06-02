package auth

import (
	"database/sql"

	"github.com/Insid1/go-auth-user/user-service/internal/model"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Login(data *model.Login, passHash []byte) (string, error) {
	var usr model.User

	row := r.DB.QueryRow("SELECT * FROM \"user\" WHERE email = $1 AND password=$2", data.Email, passHash)
	err := row.Scan(&usr)
	if err != nil {
		return "", err
	}

	return "", nil
}
