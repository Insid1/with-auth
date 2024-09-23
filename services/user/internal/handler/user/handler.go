package user

import (
	"context"

	"github.com/Insid1/go-auth-user/pkg/grpc/user_v1"
	"github.com/Insid1/go-auth-user/user/internal/converter"
	"github.com/Insid1/go-auth-user/user/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	user_v1.UnimplementedUserV1Server
	Service service.User
}

func (h *Handler) Get(ctx context.Context, req *user_v1.GetReq) (*user_v1.GetRes, error) {
	usr, err := h.Service.Get(req.GetId(), req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %s", err.Error())
	}

	if len(usr.ID) == 0 && len(usr.Email) == 0 {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &user_v1.GetRes{
		User: converter.ToUserFromModel(usr),
	}, nil
}

func (h *Handler) Create(ctx context.Context, req *user_v1.CreateReq) (*user_v1.CreateRes, error) {
	usr := converter.ToModelFromUser(req.GetUser())

	createdUsr, err := h.Service.Create(usr, req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %s", err.Error())
	}

	return &user_v1.CreateRes{User: converter.ToUserFromModel(createdUsr)}, nil
}

func (h *Handler) Update(ctx context.Context, req *user_v1.UpdateReq) (*user_v1.UpdateRes, error) {
	usr := converter.ToModelFromUser(req.GetUser())

	updatedUsr, err := h.Service.Update(usr, req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %s", err.Error())
	}

	return &user_v1.UpdateRes{User: converter.ToUserFromModel(updatedUsr)}, nil
}

func (h *Handler) Delete(ctx context.Context, req *user_v1.DeleteReq) (*user_v1.DeleteRes, error) {

	id, err := h.Service.Delete(req.GetId())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error: %s", err.Error())
	}

	return &user_v1.DeleteRes{Id: id}, nil
}

func (h *Handler) CheckPassword(ctx context.Context, req *user_v1.CheckPasswordReq) (*user_v1.CheckPasswordRes, error) {
	usr, err := h.Service.CheckPassword(req.GetId(), req.GetEmail(), req.GetPassword())

	if err != nil {
		return &user_v1.CheckPasswordRes{
			Success: false,
			User:    nil,
		}, status.Errorf(codes.InvalidArgument, "Invalid Data provided. Error: %s", err.Error())
	}

	return &user_v1.CheckPasswordRes{
		Success: true,
		User:    converter.ToUserFromModel(usr),
	}, nil
}
