package service

import "books-service/internal/repository"

type BooksService interface {
	CreateBook(title, author string) error
	DeleteBook(id int64) error
	GetAllBooks() ([]repository.Book, error)
	GetBook(id int64) (*repository.Book, error)
}

type booksService struct {
	repo repository.BooksRepository
}

func NewBooksService(repo repository.BooksRepository) BooksService {
	return &booksService{repo: repo}
}

func (s *booksService) CreateBook(titele, author string) error {
	book := &repository.Book{
		Title:  titele,
		Author: author,
	}
	return s.repo.CreateBook(book)
}

func (s *booksService) DeleteBook(id int64) error {
	return s.repo.DeleteBook(id)
}

func (s *booksService) GetAllBooks() ([]repository.Book, error) {
	return s.repo.GetAllBooks()
}

func (s *booksService) GetBook(id int64) (*repository.Book, error) {
	return s.repo.GetBook(id)
}
