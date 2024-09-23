package user

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/internal/common"

	"github.com/Insid1/go-auth-user/pkg/grpc/user_v1"
)

type Repository struct {
	UserClient *common.GRPCClient[user_v1.UserV1Client]
}

func (r *Repository) Get(
	ctx context.Context,
	userID string,
	email string,
) (*user_v1.User, error) {
	resp, err := r.UserClient.Client.Get(ctx, &user_v1.GetReq{Id: userID, Email: email})
	if err != nil {
		return nil, err
	}

	return resp.GetUser(), nil
}

func (r *Repository) Create(
	ctx context.Context,
	email string,
	password string,
) (*user_v1.User, error) {
	resp, err := r.UserClient.Client.Create(ctx, &user_v1.CreateReq{
		User: &user_v1.User{
			Email:    email,
			Username: email,
		},
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return resp.GetUser(), nil
}
func (r *Repository) CheckPassword(ctx context.Context, email string, password string) (*user_v1.User, error) {
	resp, err := r.UserClient.Client.CheckPassword(ctx, &user_v1.CheckPasswordReq{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return resp.GetUser(), nil
}
