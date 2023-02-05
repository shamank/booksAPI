package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/models"
	"net/http"
	"strconv"
)

func (h *Handler) getAllAuthors(c *gin.Context) {
	authors, err := h.services.AuthorItem.GetAllAuthors()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllAuthorsResponse{
		Data: authors,
	})
}

func (h *Handler) addAuthor(c *gin.Context) {

	var input models.Author

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idAuthor, err := h.services.AuthorItem.CreateAuthor(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": idAuthor,
	})

}

func (h *Handler) getAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	author, err := h.services.AuthorItem.GetAuthor(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *Handler) updateAuthor(c *gin.Context) {
	var input models.UpdateAuthorInput

	if err := c.BindJSON(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	authorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	if err := h.services.AuthorItem.UpdateAuthor(authorID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	if err := h.services.AuthorItem.DeleteAuthor(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
