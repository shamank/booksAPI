package service

import (
	"errors"
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
	return s.repo.GetAllBooks()
}

func (s *BookService) GetBook(bookID int) (models.Book, error) {

	return s.repo.GetBook(bookID)
}

func (s *BookService) CreateBook(book models.Book) (int, error) {
	bookID, err := s.repo.CreateBook(book)
	if err != nil {
		return 0, err
	}

	return bookID, nil
}

func (s *BookService) DeleteBook(bookID int) error {
	return s.repo.DeleteBook(bookID)
}

func (s *BookService) UpdateBook(bookID int, input models.UpdateBookInput) error {
	if ok := input.Validate(); !ok {
		return errors.New("no changes")
	}

	return s.repo.UpdateBook(bookID, input)
}
