package models

type Book struct {
	ID          int      `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Rating      *float32 `json:"user_rating"`
}

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type BookAuthor struct {
	ID       int `json:"id"`
	BookID   int `json:"book_id"`
	AuthorID int `json:"author_id"`
}

type UpdateBookInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i *UpdateBookInput) Validate() bool {
	return i.Title != nil || i.Description != nil
}
