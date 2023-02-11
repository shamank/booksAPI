package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shamank/booksAPI/models"
	"strings"
)

type AuthorPostgres struct {
	db *sqlx.DB
}

func NewAuthorPostgres(db *sqlx.DB) *AuthorPostgres {
	return &AuthorPostgres{
		db: db,
	}
}

func (r *AuthorPostgres) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author

	query := fmt.Sprintf("SELECT * FROM %s", authorsTable)

	if err := r.db.Select(&authors, query); err != nil {
		return authors, err
	}

	return authors, nil

}

func (r *AuthorPostgres) GetAuthor(authorID int) (models.Author, error) {
	var author models.Author

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", authorsTable)

	if err := r.db.Get(&author, query, authorID); err != nil {
		return author, err
	}
	return author, nil
}

func (r *AuthorPostgres) CreateAuthor(author models.Author) (int, error) {

	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s(first_name, last_name) VALUES ($1, $2) RETURNING id", authorsTable)

	row := r.db.QueryRow(query, author.FirstName, author.LastName)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *AuthorPostgres) DeleteAuthor(authorID int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", authorsTable)

	if _, err := r.db.Exec(query, authorID); err != nil {
		return err
	}
	return nil
}

func (r *AuthorPostgres) UpdateAuthor(authorID int, input models.UpdateAuthorInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", argID))
		args = append(args, input.FirstName)
		argID++
	}
	if input.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argID))
		args = append(args, input.LastName)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", authorsTable, setQuery, argID)

	args = append(args, authorID)

	_, err := r.db.Exec(query, args...)

	return err

}
