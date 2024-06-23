package handler

import (
	"context"

	"github.com/Insid1/go-auth-user/user/internal/handler/user"
	"github.com/Insid1/go-auth-user/user/internal/service"
	"github.com/Insid1/go-auth-user/user/pkg/user_v1"
)

type User interface {
	user_v1.UserV1Server
	Get(context.Context, *user_v1.GetRequest) (*user_v1.GetResponse, error)
	Create(context.Context, *user_v1.CreateRequest) (*user_v1.CreateResponse, error)
}

func NewUserHandler(ctx context.Context, srvc service.User) User {
	return &user.Handler{
		Service: srvc,
	}
}
