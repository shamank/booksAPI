package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/models"
	"net/http"
	"strconv"
)

//	@Summary		getAllBooks
//	@Description	Get All Books from DB
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	getAllBooksResponse	"books"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/books/ [get]
func (h *Handler) getAllBooks(c *gin.Context) {

	books, err := h.services.BookItem.GetAllBooks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBooksResponse{
		Books: books,
	})
}

//	@Summary		addBook
//	@Description	Create new Book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			input	body	models.Book	true	"book input"
//	@Security		ApiKeyAuth
//	@Success		200	{int}		int	id	"new book id"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/books/ [post]
func (h *Handler) createBook(c *gin.Context) {

	var inputBook models.Book

	if err := c.BindJSON(&inputBook); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idBook, err := h.services.BookItem.CreateBook(inputBook)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": idBook,
	})
}

//	@Summary		getBook
//	@Description	Get Book by ID
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"book id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	models.Book	"book"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/books/{id} [get]
func (h *Handler) getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	book, err := h.services.GetBook(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

//	@Summary		updateBook
//	@Description	Update Book by ID
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int						true	"book id"
//	@Param			input	body	models.UpdateBookInput	true	"params to update"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/books/{id} [put]
func (h *Handler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	var input models.UpdateBookInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err = h.services.BookItem.UpdateBook(id, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

//	@Summary		deleteBook
//	@Description	Delete Book by ID
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"book id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/books/{id} [delete]
func (h *Handler) deleteBook(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	err = h.services.BookItem.DeleteBook(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
