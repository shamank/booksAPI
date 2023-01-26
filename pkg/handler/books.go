package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/models"
	"net/http"
)

func (h *Handler) getAllBooks(c *gin.Context) {

}

func (h *Handler) createBook(c *gin.Context) {
	_, roleID, err := getUser(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("ROLEID:", roleID)

	if roleID == 0 {
		newErrorResponse(c, http.StatusForbidden, "you are not moderator/admin")
		return
	}

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

func (h *Handler) getBook(c *gin.Context) {

}

func (h *Handler) updateBook(c *gin.Context) {

}

func (h *Handler) deleteBook(c *gin.Context) {

}
