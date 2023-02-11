package models

type Book struct {
	ID          int      `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Rating      *float32 `json:"user_rating"`
	Author      Author   `json:"author"`
	AuthorID    int      `json:"author_id" db:"author_id"`
}

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Birth     string `json:"birth" db:"birth"`
}

type BookAuthor struct {
	ID       int `json:"id"`
	BookID   int `json:"book_id"`
	AuthorID int `json:"author_id"`
}

type UpdateBookInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	AuthorID    *int    `json:"author_id"`
}

func (i *UpdateBookInput) Validate() bool {
	return i.Title != nil || i.Description != nil || i.AuthorID != nil
}

type UpdateAuthorInput struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

func (i *UpdateAuthorInput) Validate() bool {
	return i.FirstName != nil || i.LastName != nil
}
