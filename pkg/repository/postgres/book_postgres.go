package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shamank/booksAPI/models"
	"strings"
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

	if err := r.db.Select(&books, query); err != nil {
		return books, err
	}

	return books, nil
}

func (r *BookPostgres) GetBook(bookID int) (models.Book, error) {
	var book models.Book

	query := fmt.Sprintf(`SELECT *, ((SELECT AVG(user_rating)
               from user_book WHERE book_id = $1 GROUP BY book_id)) as rating 
				FROM %s WHERE id = $1`, booksTable)
	if err := r.db.Get(&book, query, bookID); err != nil {
		return book, err
	}

	return book, nil

}

func (r *BookPostgres) CreateBook(book models.Book) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	createBookQuery := fmt.Sprintf("INSERT INTO %s(title, description, author_id) VALUES($1, $2, $3) RETURNING id", booksTable)
	row := tx.QueryRow(createBookQuery, book.Title, book.Description, book.Author.ID)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *BookPostgres) DeleteBook(bookID int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", booksTable)

	_, err := r.db.Exec(query, bookID)
	return err

}

func (r *BookPostgres) UpdateBook(bookID int, input models.UpdateBookInput) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.AuthorID != nil {
		setValues = append(setValues, fmt.Sprintf("author_id=$%d", argId))
		args = append(args, *input.AuthorID)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d",
		booksTable, setQuery, argId)
	args = append(args, bookID)

	_, err := r.db.Exec(query, args...)
	return err

}
