package service

import (
	"errors"
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

func (s *UserService) GetUserBooks(userID int) ([]models.Book, error) {

	return s.repo.GetUserBooks(userID)
}

func (s *UserService) GetUserAuthors(userID int) ([]models.Author, error) {

	return s.repo.GetUserAuthors(userID)
}

func (s *UserService) NewUserBook(userID int, bookID int) error {

	return s.repo.NewUserBook(userID, bookID)
}

func (s *UserService) NewUserAuthor(userID int, authorID int) error {

	return s.repo.NewUserAuthor(userID, authorID)
}

func (s *UserService) RemoveUserBook(userID int, bookID int) error {

	return s.repo.RemoveUserBook(userID, bookID)
}

func (s *UserService) RemoveUserAuthor(userID int, authorID int) error {
	return s.repo.RemoveUserAuthor(userID, authorID)
}

func (s *UserService) UpdateUser(userID int, input models.UserUpdate) error {

	if ok := input.Validate(); !ok {
		return errors.New("nothing to update")
	}

	return s.repo.UpdateUser(userID, input)
}
