package repository

import "gorm.io/gorm"

type BookManagement struct {
	ID     int64 `gorm:"primaryKey"`
	UserID int64 `gorm:"not null"`
	BookID int64 `gorm:"not null"`
}

type BookManagementRepository interface {
	CheckOutBook(userID,bookID int64) error
	CheckInBook(userID,bookID int64) error
}

type bookManagementRepository struct {
	db *gorm.DB
}

func NewBookManagementRepository(db *gorm.DB) BookManagementRepository {
    return &bookManagementRepository{db: db}
}

func (r *bookManagementRepository) CheckOutBook(userID,bookID int64)error{
	bm:=&BookManagement{
		BookID: bookID,
		UserID: userID,
	}
	return r.db.Create(bm).Error
}

func (r *bookManagementRepository) CheckInBook(userID,bookID int64) error{
	return r.db.Where("user_id = ? AND book_id = ?",userID,bookID).Delete(&BookManagement{}).Error
}