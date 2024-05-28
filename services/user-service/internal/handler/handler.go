package handler

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/internal/handler/auth"
	"github.com/Insid1/go-auth-user/user-service/internal/handler/user"
	"github.com/Insid1/go-auth-user/user-service/internal/service"
	"github.com/Insid1/go-auth-user/user-service/pkg/auth_v1"
	"github.com/Insid1/go-auth-user/user-service/pkg/user_v1"
)

type User interface {
	user_v1.UserV1Server
	Get(context.Context, *user_v1.GetRequest) (*user_v1.GetResponse, error)
	Create(context.Context, *user_v1.CreateRequest) (*user_v1.CreateResponse, error)
}

type Auth interface {
	auth_v1.AuthV1Server
	Login(context.Context, *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error)
	Register(context.Context, *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error)
	Logout(context.Context, *auth_v1.LogoutRequest) (*auth_v1.LogoutResponse, error)
}

func NewUserHandler(ctx context.Context, srvc service.User) User {
	return &user.Handler{
		Service: srvc,
	}
}

func NewAuthHandler(ctx context.Context) Auth {
	return &auth.Handler{}
}
