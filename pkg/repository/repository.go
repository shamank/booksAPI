package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/shamank/booksAPI/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type AuthorItem interface {
}

type BookItem interface {
	GetAllBooks() ([]models.Book, error)
	GetBook(bookID int) (models.Book, error)

	CreateBook(book models.Book) (int, error)
}

type Repository struct {
	Authorization
	AuthorItem
	BookItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		BookItem:      NewBookPostgres(db),
	}
}
