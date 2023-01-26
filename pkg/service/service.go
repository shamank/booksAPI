package service

import (
	"github.com/shamank/booksAPI/models"
	"github.com/shamank/booksAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, int, error)
}

type AuthorItem interface {
}

type BookItem interface {
	GetAllBooks() ([]models.Book, error)
	GetBook(bookID int) (models.Book, error)

	CreateBook(book models.Book) (int, error)
}

type Service struct {
	Authorization
	AuthorItem
	BookItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		BookItem:      NewBookService(repos.BookItem),
	}
}
