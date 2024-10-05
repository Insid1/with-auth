package user

import (
	"context"

	"github.com/Insid1/with-auth/pkg/grpc/user_v1"
	userPkg "github.com/Insid1/with-auth/user/pkg"
)

type Repository struct {
	UserClient *userPkg.GRPCInitializedUserClient
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
