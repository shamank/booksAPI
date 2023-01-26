package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shamank/booksAPI/models"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) GetAllBooks() ([]models.Book, error) {

	var books []models.Book

	query := fmt.Sprintf("SELECT * FROM %s", booksTable)

	if err := r.db.Get(&books, query); err != nil {
		return books, err
	}

	return books, nil
}

func (r *BookPostgres) GetBook(bookID int) (models.Book, error) {
	var book models.Book

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", booksTable)
	if err := r.db.Get(&book, query, bookID); err != nil {
		return book, err
	}

	//result, err := r.db.QueryRow(query, bookID)
	//if err != nil {
	//	return book, err
	//}
	//if err := result.Scan(&book); err != nil {
	//	return book, err
	//}

	return book, nil

}

func (r *BookPostgres) CreateBook(book models.Book) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	createBookQuery := fmt.Sprintf("INSERT INTO %s(title, description, rating) VALUES($1, $2, $3) RETURNING id", booksTable)
	row := tx.QueryRow(createBookQuery, book.Title, book.Description, book.Rating)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
