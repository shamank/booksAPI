package service

import (
	"errors"
	"github.com/shamank/booksAPI/models"
	"github.com/shamank/booksAPI/pkg/repository"
)

type AuthorService struct {
	repo repository.AuthorItem
}

func NewAuthorService(repo repository.AuthorItem) *AuthorService {
	return &AuthorService{
		repo: repo,
	}
}

func (s *AuthorService) GetAllAuthors() ([]models.Author, error) {

	return s.repo.GetAllAuthors()
}

func (s *AuthorService) GetAuthor(authorID int) (models.Author, error) {

	return s.repo.GetAuthor(authorID)

}

func (s *AuthorService) CreateAuthor(author models.Author) (int, error) {

	authorID, err := s.repo.CreateAuthor(author)
	if err != nil {
		return 0, err
	}

	return authorID, nil
}

func (s *AuthorService) DeleteAuthor(authorID int) error {

	return s.repo.DeleteAuthor(authorID)
}

func (s *AuthorService) UpdateAuthor(authorID int, input models.UpdateAuthorInput) error {

	if ok := input.Validate(); ok {
		return errors.New("nothing to change")
	}
	return s.repo.UpdateAuthor(authorID, input)
}