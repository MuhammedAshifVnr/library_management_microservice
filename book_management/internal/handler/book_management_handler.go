package handler

import (
	"book_management_service/internal/service"
	pb "book_management_service/proto"
	"context"
)

type BookManagementHandler struct {
	pb.UnimplementedBookManagementServiceServer
	service service.BookManagementService
}

func NewBookManagementHandler(service service.BookManagementService) *BookManagementHandler {
	return &BookManagementHandler{service: service}
}

func (h *BookManagementHandler) CheckOutBook(ctx context.Context, req *pb.CheckOutBookRequest) (*pb.CheckOutBookResponse, error) {
	err := h.service.CheckOutBook(req.UserId, req.BookId)
	if err != nil {
		return nil, err
	}
	return &pb.CheckOutBookResponse{Success: true}, nil
}

func (h *BookManagementHandler) CheckInBook(ctx context.Context, req *pb.CheckInBookRequest) (*pb.CheckInBookResponse, error) {
	err := h.service.CheckInBook(req.UserId, req.BookId)
	if err != nil {
		return nil, err
	}
	return &pb.CheckInBookResponse{Success: true}, nil
}
