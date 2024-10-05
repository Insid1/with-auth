package user

import (
	"database/sql"
	"fmt"

	dbErrors "github.com/Insid1/with-auth/pkg/errors/db"
	"github.com/Insid1/with-auth/user/internal/model"
	"github.com/pkg/errors"
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

	err := r.checkValidColumns([]string{column})
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	SELECT id, username, email, password_hash, created_at, updated_at FROM "users" WHERE %s=$1
	`, column)

	err = r.DB.QueryRow(
		query,
		source,
	).Scan(
		&usr.ID,
		&usr.Username,
		&usr.Email,
		&usr.PassHash,
		&usr.CreatedAt,
		&usr.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *Repository) Create(usr *model.User) (*model.User, error) {
	var createdUsr model.User

	q := fmt.Sprintf(`
		INSERT INTO "users" (username, email, password_hash) VALUES ($1, $2, $3) RETURNING %s;
		`,
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
		"UPDATE \"users\" SET %s WHERE id=$1 RETURNING %s;",
		updateStr,
		r.getReturningDBFields(),
	)

	err = r.DB.QueryRow(q, updateWith.ID).Scan(r.getReturningStructFields(&updatedUsr)...)
	if err != nil {
		return nil, err
	}

	return &updatedUsr, nil
}

func (r *Repository) Delete(userID string) (string, error) {
	_, err := r.DB.Exec("DELETE FROM \"users\" WHERE id=$1", userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (r *Repository) getReturningDBFields() string {
	return ("id, username, email, password_hash, created_at, updated_at")
}

func (r *Repository) getReturningStructFields(usr *model.User) []interface{} {
	return []interface{}{&usr.ID, &usr.Username, &usr.Email, &usr.PassHash, &usr.CreatedAt, &usr.UpdatedAt}
}

func (r *Repository) checkValidColumns(columnsToCheck []string) error {
	validColumns := map[string]bool{
		"id":            true,
		"username":      true,
		"email":         true,
		"password_hash": true,
		"created_at":    true,
		"updated_at":    true,
	}

	for _, columnToCheck := range columnsToCheck {
		_, ok := validColumns[columnToCheck]
		// Проверяем, является ли указанная колонка допустимой
		if !ok {
			return errors.Wrap(dbErrors.ErrInvalidColumn, columnToCheck)
		}
	}

	return nil
}
