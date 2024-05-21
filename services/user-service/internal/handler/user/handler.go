package user

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/pkg/user_v1"
)

type Handler struct {
	user_v1.UnimplementedUserV1Server
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Get(context.Context, *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	return nil, nil
}
