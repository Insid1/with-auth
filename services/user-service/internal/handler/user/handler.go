package user

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/internal/converter"
	"github.com/Insid1/go-auth-user/user-service/internal/service"
	"github.com/Insid1/go-auth-user/user-service/pkg/user_v1"
)

type Handler struct {
	user_v1.UnimplementedUserV1Server
	*service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		Service: svc,
	}
}

func (h *Handler) Get(ctx context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	usr := h.Service.UserService.Get(req.GetId())

	u := converter.ToUserFromModel(usr)
	return &user_v1.GetResponse{User: u}, nil
}
