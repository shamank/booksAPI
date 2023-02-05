package service

import (
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

}

func (s *AuthorService) GetAuthor(authorID int) (models.Author, error) {

}

func (s *AuthorService) CreateAuthor(author models.Author) (int, error) {

}

func (s *AuthorService) DeleteAuthor(authorID int) error {

}

func (s *AuthorService) UpdateAuthor(authorID int, input models.UpdateAuthorInput) error {

}