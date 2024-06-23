package repository

import (
	"context"
	"database/sql"

	"github.com/Insid1/go-auth-user/auth-service/internal/common"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository/auth"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository/user"

	"github.com/Insid1/go-auth-user/user/pkg/user_v1"
)

type User interface {
	Get(userID string, email string) (*user_v1.User, error)
	Create(email string, passHash string) (string, error)
}

type Auth interface {
	SaveToken(token string, userID string) error
	IsTokenLinkedWithUser(token string, userID string) bool
	RemoveToken(token string) bool
}

func NewUserRepository(ctx context.Context, client *common.GRPCClient[user_v1.UserV1Client]) User {
	return &user.Repository{
		Ctx: ctx, UserClient: client,
	}
}

func NewAuthRepository(ctx context.Context, db *sql.DB) Auth {
	return &auth.Repository{
		Ctx: ctx,
		DB:  db,
	}
}
