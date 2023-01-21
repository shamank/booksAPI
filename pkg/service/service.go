package service

import (
	"github.com/shamank/booksAPI/models"
	"github.com/shamank/booksAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type AuthorItem interface {
}

type BookItem interface {
}

type Service struct {
	Authorization
	AuthorItem
	BookItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
