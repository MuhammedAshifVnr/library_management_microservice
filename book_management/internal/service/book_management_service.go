package service

import (
	"book_management_service/internal/repository"
	booksPb "github.com/MuhammedAshifVnr/library_management_microservice"
)
type BookManagementService interface{

}

type bookManagementService struct{
	repo repository.BookManagementRepository
	booksClient booksPb.
}