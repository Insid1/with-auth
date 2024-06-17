package user

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/internal/common"
	"github.com/Insid1/go-auth-user/auth-service/pkg/user_v1"
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
