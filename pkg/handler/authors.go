package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/models"
	"net/http"
	"strconv"
)

//	@Summary		getAllAuthors
//	@Description	Get All Authors from DB
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	getAllAuthorsResponse	"authors"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/authors/ [get]
func (h *Handler) getAllAuthors(c *gin.Context) {
	authors, err := h.services.AuthorItem.GetAllAuthors()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllAuthorsResponse{
		Authors: authors,
	})
}

//	@Summary		addAuthor
//	@Description	Create new Author
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			input	body	models.Author	true	"author input"
//	@Security		ApiKeyAuth
//	@Success		200	{int}		int	id	"new author id"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/authors/ [post]
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

//	@Summary		getAuthor
//	@Description	Get Author by ID
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"author id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	models.Author	"author"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/authors/{id} [get]
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

//	@Summary		updateAuthor
//	@Description	Update Author by ID
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int							true	"author id"
//	@Param			input	body	models.UpdateAuthorInput	true	"params to update"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/authors/{id} [put]
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

//	@Summary		deleteAuthor
//	@Description	Delete Author by ID
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"author id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/authors/{id} [delete]
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
