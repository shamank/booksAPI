package service

import (
	"github.com/shamank/booksAPI/models"
	"github.com/shamank/booksAPI/pkg/repository"
)

type UserService struct {
	repo repository.UserItem
}

func NewUserService(repo repository.UserItem) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserById(userID int) (models.User, error) {

	//var user models.User

	return s.repo.GetUserById(userID)
}

func (s *UserService) GetUsersBook(userID int) ([]models.Book, error) {

	return nil, nil
}

func (s *UserService) GetUserAuthors(userID int) ([]models.Author, error) {

	return nil, nil
}

func (s *UserService) NewUserBook(userID int, bookID int) error {
	return nil
}

func (s *UserService) NewUserAuthor(userID int, authorID int) error {

	return nil
}

func (s *UserService) UpdateUser(userID int) error {
	return nil
}
