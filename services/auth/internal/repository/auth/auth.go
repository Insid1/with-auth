package auth

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) GetJWTUserKey(ctx context.Context, userID string) (string, error) {
	var jwtKey string

	err := r.DB.QueryRow("SELECT jwt_key FROM auth WHERE user_id = $1", userID).Scan(&jwtKey)

	if err != nil {
		return "", err
	}

	return jwtKey, nil
}

func (r *Repository) GenerateJWTUserKey(ctx context.Context, userID string) (string, error) {
	var jwtKey string
	err := r.DB.QueryRow("INSERT INTO auth (user_id) VALUES ($1) RETURNING jwt_key;", userID).Scan(&jwtKey)

	if err != nil {
		return "", err
	}

	return jwtKey, nil
}
