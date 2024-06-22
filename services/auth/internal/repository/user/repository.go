package user

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/internal/common"

	"github.com/Insid1/go-auth-user/user-service/pkg/user_v1"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Repository struct {
	Ctx        context.Context
	UserClient *common.GRPCClient[user_v1.UserV1Client]
}

func (r *Repository) Get(userID string, email string) (*user_v1.User, error) {
	resp, err := r.UserClient.Client.Get(r.Ctx, &user_v1.GetRequest{Id: userID, Email: email})
	if err != nil {
		return nil, err
	}

	return resp.GetUser(), nil
}

func (r *Repository) Create(email string, passHash string) (string, error) {
	resp, err := r.UserClient.Client.Create(r.Ctx, &user_v1.CreateRequest{
		User: &user_v1.User{
			Email:    email,
			Age:      0,
			Name:     &wrapperspb.StringValue{},
			PassHash: passHash,
		},
	})
	if err != nil {
		return "", err
	}

	return resp.GetId(), nil
}
