package user

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/internal/converter"
	"github.com/Insid1/go-auth-user/user-service/internal/service"
	"github.com/Insid1/go-auth-user/user-service/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	usr, err := h.Service.UserService.Get(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %s", err.Error())
	}

	if len(usr.ID) == 0 {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &user_v1.GetResponse{
		User: converter.ToUserFromModel(usr),
	}, nil
}

func (h *Handler) Create(ctx context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	usr := converter.ToModelFromUser(req.GetUser())

	id, err := h.UserService.Create(usr)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %s", err.Error())
	}

	return &user_v1.CreateResponse{Id: id}, nil
}
