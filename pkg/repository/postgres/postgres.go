package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable   = "users"
	booksTable   = "books"
	authorsTable = "authors"

	userAuthorTable = "user_author"
	userBookTable   = "user_book"

	bookAuthor = "book_author"
	rolesTable = "usr_roles"
)

type ConfigDB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg ConfigDB) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
