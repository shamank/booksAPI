package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/shamank/booksAPI/models"
	"github.com/shamank/booksAPI/pkg/repository/postgres"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
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
	GetUserBooks(userID int) ([]models.Book, error)
	GetUserAuthors(userID int) ([]models.Author, error)

	NewUserBook(userID int, bookID int) error
	NewUserAuthor(userID int, authorID int) error

	RemoveUserBook(userID int, bookID int) error
	RemoveUserAuthor(userID int, bookID int) error

	SetBookRating(userID int, bookID int, input models.UserScoreBook) error

	UpdateUser(userID int, input models.UserUpdate) error
}

type Repository struct {
	Authorization
	AuthorItem
	BookItem
	UserItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		BookItem:      postgres.NewBookPostgres(db),
		AuthorItem:    postgres.NewAuthorPostgres(db),
		UserItem:      postgres.NewUserPostgres(db),
	}
}
