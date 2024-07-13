package handler

import (
	"github.com/MuhammedAshifVnr/library_management_microservice/books/internal/service"
	pb "github.com/MuhammedAshifVnr/library_management_microservice/books/proto"
	"golang.org/x/net/context"
)

type BooksHandler struct {
	pb.UnimplementedBookServiceServer
	service service.BooksService
}

func NewBooksHandler(service service.BooksService) *BooksHandler {
	return &BooksHandler{service: service}
}

func (h *BooksHandler) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	err := h.service.CreateBook(req.Title, req.Author)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBookResponse{Success: true}, nil
}

func (h *BooksHandler) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	err := h.service.DeleteBook(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBookResponse{Success: true}, nil
}

func (h *BooksHandler) GetAllBooks(ctx context.Context, req *pb.GetAllBooksRequest) (*pb.GetAllBooksResponse, error) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		return nil, err
	}
	var pbBooks []*pb.Book
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:     book.ID,
			Title:  book.Title,
			Author: book.Author,
		})
	}
	return &pb.GetAllBooksResponse{Books: pbBooks}, nil

}

func (h *BooksHandler) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	book, err := h.service.GetBook(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetBookResponse{Book: &pb.Book{
		Id:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}}, nil
}
