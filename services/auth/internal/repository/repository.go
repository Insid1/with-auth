package repository

import (
	"context"
	"database/sql"

	"github.com/Insid1/go-auth-user/auth-service/internal/common"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository/auth"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository/user"
	"github.com/Insid1/go-auth-user/pkg/grpc/user_v1"
)

type User interface {
	Get(ctx context.Context, userID string, email string) (*user_v1.User, error)
	CheckPassword(ctx context.Context, email string, password string) (*user_v1.User, error)
	Create(ctx context.Context, email string, password string) (*user_v1.User, error)
}

type Auth interface {
	SaveToken(ctx context.Context, token string, userID string) error
	IsTokenLinkedWithUser(ctx context.Context, token string, userID string) bool
	RemoveToken(ctx context.Context, token string) bool
}

func NewUserRepository(client *common.GRPCClient[user_v1.UserV1Client]) User {
	return &user.Repository{
		UserClient: client,
	}
}

func NewAuthRepository(db *sql.DB) Auth {
	return &auth.Repository{
		DB: db,
	}
}
