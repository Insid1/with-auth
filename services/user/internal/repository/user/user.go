package user

import (
	"database/sql"
	"fmt"

	"github.com/Insid1/with-auth/user/internal/model"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Get(id string) (*model.User, error) {
	var usr model.User

	err := r.DB.QueryRow(
		"SELECT id, username, email, password_hash, created_at, updated_at FROM \"users\" WHERE id=$1", id).Scan(
		&usr.ID, &usr.Username, &usr.Email, &usr.PassHash, &usr.CreatedAt, &usr.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *Repository) GetBy(column string, source string) (*model.User, error) {
	var usr model.User
	query := fmt.Sprintf("SELECT id, username, email, password_hash, created_at, updated_at FROM \"users\" WHERE %s=$1", column)

	err := r.DB.QueryRow(query, source).Scan(&usr.ID, &usr.Username, &usr.Email, &usr.PassHash, &usr.CreatedAt, &usr.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *Repository) Create(usr *model.User) (*model.User, error) {
	var createdUsr model.User

	q := fmt.Sprintf(
		"INSERT INTO %s (username, email, password_hash) VALUES ($1, $2, $3) RETURNING %s;",
		r.getUsersDBName(),
		r.getReturningDBFields(),
	)

	err := r.DB.QueryRow(
		q,
		usr.Username,
		usr.Email,
		usr.PassHash,
	).Scan(r.getReturningStructFields(&createdUsr)...)

	if err != nil {
		return nil, err
	}

	return &createdUsr, nil
}

func (r *Repository) Update(updateWith *model.User) (*model.User, error) {
	var updatedUsr model.User

	updateStr, err := updateWith.BuildUpdateString()
	if err != nil {
		return nil, err
	}

	q := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id=$1 RETURNING %s;",
		r.getUsersDBName(),
		updateStr,
		r.getReturningDBFields(),
	)

	err = r.DB.QueryRow(q, updateWith.ID).Scan(r.getReturningStructFields(&updatedUsr)...)

	if err != nil {
		return nil, err
	}

	return &updatedUsr, nil
}

func (r *Repository) Delete(id string) (string, error) {
	q := fmt.Sprintf(
		"DELETE FROM %s WHERE id=$1",
		r.getUsersDBName(),
	)
	_, err := r.DB.Exec(q, id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *Repository) getReturningDBFields() string {
	return ("id, username, email, password_hash, created_at, updated_at")
}

func (r *Repository) getReturningStructFields(usr *model.User) []interface{} {
	return []interface{}{&usr.ID, &usr.Username, &usr.Email, &usr.PassHash, &usr.CreatedAt, &usr.UpdatedAt}
}

func (r *Repository) getUsersDBName() string {

	return "\"users\""
}
