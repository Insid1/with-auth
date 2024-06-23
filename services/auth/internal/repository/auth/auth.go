package auth

import (
	"context"
	"database/sql"
)

type Repository struct {
	Ctx context.Context
	DB  *sql.DB
}

func (r *Repository) SaveToken(token string, userID string) error {
	_, err := r.DB.Exec("INSERT INTO token (token, user_id) VALUES ($1, $2)", token, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) IsTokenLinkedWithUser(token string, userID string) bool {
	var tokenId string

	err := r.DB.QueryRow("SELECT id FROM token WHERE token = $1 AND user_id = $2", token, userID).Scan(&tokenId)

	return err == nil
}

func (r *Repository) RemoveToken(token string) bool {
	_, err := r.DB.Exec("DELETE FROM token WHERE token = $1", token)

	return err == nil
}
