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
	GetAllAuthors() ([]models.Author, error)
	GetAuthor(authorID int) (models.Author, error)

	CreateAuthor(author models.Author) (int, error)
	DeleteAuthor(authorID int) error
	UpdateAuthor(authorID int, input models.UpdateAuthorInput) error
}

type BookItem interface {
	GetAllBooks() ([]models.Book, error)
	GetBook(bookID int) (models.Book, error)

	CreateBook(book models.Book) (int, error)
	DeleteBook(bookID int) error
	UpdateBook(bookID int, input models.UpdateBookInput) error
}

type UserItem interface {
	GetUserById(userID int) (models.User, error)
	GetUsersBook(userID int) ([]models.Book, error)
	GetUserAuthors(userID int) ([]models.Author, error)
	NewUserBook(userID int, bookID int) error
	NewUserAuthor(userID int, authorID int) error
	UpdateUser(userID int) error
}

type Service struct {
	Authorization
	AuthorItem
	BookItem
	UserItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		BookItem:      NewBookService(repos.BookItem),
		AuthorItem:    NewAuthorService(repos.AuthorItem),
		UserItem:      NewUserService(repos.UserItem),
	}
}
