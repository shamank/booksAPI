package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shamank/booksAPI/models"
	"strings"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) GetUserById(userID int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT ut.id, name, username, role_id, rt.title as role_name FROM %s ut INNER JOIN %s rt ON ut.role_id = rt.id  WHERE ut.id=$1", usersTable, rolesTable)

	err := r.db.Get(&user, query, userID)

	return user, err
}

func (r *UserPostgres) GetUserBooks(userID int) ([]models.Book, error) {

	var books []models.Book

	query := fmt.Sprintf("SELECT bt.* FROM %s bt INNER JOIN %s ubt ON bt.id = ubt.book_id WHERE ubt.user_id = $1", booksTable, userBookTable)

	if err := r.db.Select(&books, query, userID); err != nil {
		return books, err
	}

	return books, nil

}

func (r *UserPostgres) GetUserAuthors(userID int) ([]models.Author, error) {

	var authors []models.Author

	query := fmt.Sprintf("SELECT bt.* FROM %s bt INNER JOIN %s uat ON bt.id = uat.author_id WHERE uat.user_id = $1", authorsTable, userAuthorTable)

	if err := r.db.Select(&authors, query, userID); err != nil {
		return authors, err
	}

	return authors, nil

}

func (r *UserPostgres) NewUserBook(userID int, bookID int) error {
	query := fmt.Sprintf("INSERT INTO %s(user_id, book_id) VALUES ($1, $2)", userBookTable)

	if _, err := r.db.Exec(query, userID, bookID); err != nil {
		return err
	}
	return nil
}

func (r *UserPostgres) NewUserAuthor(userID int, authorID int) error {
	query := fmt.Sprintf("INSERT INTO %s(user_id, author_id) VALUES ($1, $2)", userAuthorTable)

	if _, err := r.db.Exec(query, userID, authorID); err != nil {
		return err
	}
	return nil
}

func (r *UserPostgres) RemoveUserBook(userID int, bookID int) error {

	query := fmt.Sprintf("DELETE from %s WHERE user_id = $1 and book_id = $2", userBookTable)

	_, err := r.db.Exec(query, userID, bookID)
	return err
}

func (r *UserPostgres) RemoveUserAuthor(userID int, authorID int) error {
	query := fmt.Sprintf("DELETE from %s WHERE user_id = $1 and author_id = $2", userAuthorTable)

	_, err := r.db.Exec(query, userID, authorID)

	return err
}

func (r *UserPostgres) UpdateUser(userID int, input models.UserUpdate) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *input.Username)
		argId++
	}

	if input.RoleID != nil {
		setValues = append(setValues, fmt.Sprintf("role_id=$%d", argId))
		args = append(args, *input.RoleID)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d",
		usersTable, setQuery, argId)
	args = append(args, userID)

	_, err := r.db.Exec(query, args...)
	return err

}
