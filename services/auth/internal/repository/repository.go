package repository

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/internal/common"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository/user"
	"github.com/Insid1/go-auth-user/auth-service/pkg/user_v1"
)

type User interface {
	Get(userID string, email string) (*user_v1.User, error)
}

func NewUserRepository(ctx context.Context, client *common.GRPCClient[user_v1.UserV1Client]) User {
	return &user.Repository{
		Ctx: ctx, UserClient: client,
	}
}
