package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID    int64 `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}

type UsersRepository interface {
	CreateUser(user *User) error
	DeleteUser(id int64) error
	GetAllUsers() ([]User, error)
	GetUser(id int64) (*User, error)
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db: db}
}

func (r *usersRepository) CreateUser(user *User) error {
	if user.Email == "" || user.Name == ""{
		return fmt.Errorf("all fields are requierd")
	}
	return r.db.Create(user).Error
}

func (r *usersRepository) DeleteUser(id int64) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *usersRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *usersRepository) GetUser(id int64) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}
