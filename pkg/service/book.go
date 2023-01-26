package service

import (
	"github.com/shamank/booksAPI/models"
	"github.com/shamank/booksAPI/pkg/repository"
)

type BookService struct {
	repo repository.BookItem
}

func NewBookService(repo repository.BookItem) *BookService {
	return &BookService{repo: repo}
}

//	GetAllBooks() ([]models.Book, error)
//	GetBook(bookID int) (models.Book, error)

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return nil, nil
}

func (s *BookService) GetBook(bookID int) (models.Book, error) {
	var book models.Book
	return book, nil
}

func (s *BookService) CreateBook(book models.Book) (int, error) {
	bookID, err := s.repo.CreateBook(book)
	if err != nil {
		return 0, err
	}

	return bookID, nil
}
