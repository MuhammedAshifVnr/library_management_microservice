package service

import "users-service/internal/repository"

type UsersService interface {
	CreateUser(name, email string) error
	DeleteUser(id int64) error
	GetAllUsers() ([]repository.User, error)
	GetUser(id int64) (*repository.User, error)
}

type usersService struct {
	repo repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository) UsersService {
	return &usersService{repo: repo}
}

func (s *usersService) CreateUser(name, email string) error {
	user := &repository.User{
		Name:  name,
		Email: email,
	}
	return s.repo.CreateUser(user)
}

func (s *usersService) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}

func (s *usersService) GetAllUsers() ([]repository.User, error) {
	return s.repo.GetAllUsers()
}

func (s *usersService) GetUser(id int64) (*repository.User, error) {
	return s.repo.GetUser(id)
}
