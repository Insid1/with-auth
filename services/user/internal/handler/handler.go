package handler

import (
	"context"

	"github.com/Insid1/go-auth-user/pkg/grpc/user_v1"
	"github.com/Insid1/go-auth-user/user/internal/handler/user"
	"github.com/Insid1/go-auth-user/user/internal/service"
)

type User interface {
	user_v1.UserV1Server

	Create(context.Context, *user_v1.CreateReq) (*user_v1.CreateRes, error)
	Get(context.Context, *user_v1.GetReq) (*user_v1.GetRes, error)
	Update(context.Context, *user_v1.UpdateReq) (*user_v1.UpdateRes, error)
	Delete(context.Context, *user_v1.DeleteReq) (*user_v1.DeleteRes, error)
}

func NewUserHandler(srvc service.User) User {
	return &user.Handler{
		Service: srvc,
	}
}
