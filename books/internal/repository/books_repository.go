package repository

import (
	"fmt"

	"gorm.io/gorm"
)


type Book struct {
    ID     int64  `gorm:"primaryKey"`
    Title  string   `gorm:"unique;not null"`
    Author string   `gorm:"not null"`
}

type BooksRepository interface {
    CreateBook(book *Book) error
    DeleteBook(id int64) error
    GetAllBooks() ([]Book, error)
    GetBook(id int64) (*Book, error)
}

type booksRepository struct {
    db *gorm.DB
}

func NewBooksRepository(db *gorm.DB) BooksRepository {
    return &booksRepository{db: db}
}

func (r *booksRepository) CreateBook(book *Book) error {
    if book.Author == "" || book.Title==""{
        return fmt.Errorf("all fields are required")
    }
    return r.db.Create(book).Error
}

func (r *booksRepository) DeleteBook(id int64) error {
    return r.db.Delete(&Book{}, id).Error
}

func (r *booksRepository) GetAllBooks() ([]Book, error) {
    var books []Book
    err := r.db.Find(&books).Error
    return books, err
}

func (r *booksRepository) GetBook(id int64) (*Book, error) {
    var book Book
    err := r.db.First(&book, id).Error
    return &book, err
}
