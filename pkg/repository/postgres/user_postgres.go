package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shamank/booksAPI/models"
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

func (r *UserPostgres) GetUsersBook(userID int) ([]models.Book, error) {
	return nil, nil
}

func (r *UserPostgres) GetUserAuthors(userID int) ([]models.Author, error) {
	return nil, nil
}

func (r *UserPostgres) NewUserBook(userID int, bookID int) error {
	return nil
}

func (r *UserPostgres) NewUserAuthor(userID int, authorID int) error {
	return nil
}

func (r *UserPostgres) UpdateUser(userID int) error {
	return nil
}
