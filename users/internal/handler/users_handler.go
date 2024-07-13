package handler

import (
	"context"

	"users-service/internal/service"
	pb "users-service/proto"
)

type UsersHandler struct {
	pb.UnimplementedUserServiceServer
	service service.UsersService
}

func NewUsersHandler(service service.UsersService) *UsersHandler {
	return &UsersHandler{service: service}
}

func (h *UsersHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := h.service.CreateUser(req.Name, req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Success: true}, nil
}

func (h *UsersHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := h.service.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Success: true}, nil
}

func (h *UsersHandler) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return &pb.GetAllUsersResponse{Users: pbUsers}, nil
}

func (h *UsersHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.service.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{User: &pb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}}, nil
}
