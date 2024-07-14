package service

import (
	"book_management_service/internal/repository"
	pb "book_management_service/proto"
	"context"
)

type BookManagementService interface {
	CheckOutBook(userID,bookID int64)error
	CheckInBook(userID,bookID int64)error
}

type bookManagementService struct {
	repo        repository.BookManagementRepository
	booksClient pb.BookServiceClient
	usersClient pb.UserServiceClient
}

func NewBookManagementService(repo repository.BookManagementRepository, booksClient pb.BookServiceClient, usersClient pb.UserServiceClient) BookManagementService {
    return &bookManagementService{
        repo:        repo,
        booksClient: booksClient,
        usersClient: usersClient,
    }
}

func (s *bookManagementService) CheckOutBook (userID,bookID int64)error{
	userRes,err :=s.usersClient.GetUser(context.Background(), &pb.GetUserRequest{Id: userID})
	if err !=nil{
		return err
	}

	bookRes,err :=s.booksClient.GetBook(context.Background(), &pb.GetBookRequest{Id: bookID})
	if err !=nil{
		return err
	}

	return s.repo.CheckOutBook(userRes.User.Id,bookRes.Book.Id)
}

func (s *bookManagementService) CheckInBook(userID,bookID int64)error{
	return s.repo.CheckInBook(userID,bookID)
}