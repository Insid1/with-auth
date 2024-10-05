package repository

import (
	"context"
	"database/sql"

	"github.com/Insid1/go-auth-user/auth-service/internal/repository/auth"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository/user"
	"github.com/Insid1/go-auth-user/pkg/grpc/user_v1"
	userPkg "github.com/Insid1/go-auth-user/user/pkg"
)

type User interface {
	Get(ctx context.Context, userID string, email string) (*user_v1.User, error)
	CheckPassword(ctx context.Context, email string, password string) (*user_v1.User, error)
	Create(ctx context.Context, email string, password string) (*user_v1.User, error)
}

type Auth interface {
	GetJWTUserKey(ctx context.Context, userID string) (string, error)
	GenerateJWTUserKey(ctx context.Context, userID string) (string, error)
}

func NewUserRepository(client *userPkg.GRPCInitializedUserClient) User {
	return &user.Repository{
		UserClient: client,
	}
}

func NewAuthRepository(db *sql.DB) Auth {
	return &auth.Repository{
		DB: db,
	}
}
