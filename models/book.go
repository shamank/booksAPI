package models

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
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
