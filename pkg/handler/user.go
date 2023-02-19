package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/models"
	"net/http"
	"strconv"
)

//	@Summary		getUser
//	@Description	Get User by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int	true	"user id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	models.User	"user"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/{id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	user, err := h.services.UserItem.GetUserById(id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

//	@Summary		getUserBooks
//	@Description	Get User Books list
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int	true	"user id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	getAllBooksResponse	"books"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/{user_id}/books [get]
func (h *Handler) getUserBooks(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	books, err := h.services.GetUserBooks(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBooksResponse{Books: books})
}

//	@Summary		getUserAuthors
//	@Description	Get User authors list by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int	true	"user id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	getAllAuthorsResponse	"authors"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/{user_id}/authors [get]
func (h *Handler) getUserAuthors(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	authors, err := h.services.GetUserAuthors(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllAuthorsResponse{Authors: authors})

}

//	@Summary		newUserBook
//	@Description	Add new book in list for User by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int	true	"user id"
//	@Param			book_id	path	int	true	"book id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/{user_id}/books/{book_id} [post]
func (h *Handler) newUserBook(c *gin.Context) {

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	if err := h.services.NewUserBook(userID, bookID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

//	@Summary		setBookRating
//	@Description	Set book rating for User by IDs
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int						true	"user id"
//	@Param			book_id	path	int						true	"book id"
//	@Param			book_id	body	models.UserScoreBook	true	"user score"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/{user_id}/books/{book_id} [put]
func (h *Handler) setBookRating(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: user_id")
		return
	}

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: book_id")
		return
	}

	var input models.UserScoreBook
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.SetBookRating(userID, bookID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

//	@Summary		newUserAuthor
//	@Description	Add new author in list for User by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id		path	int	true	"user id"
//	@Param			author_id	path	int	true	"author id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/{user_id}/authors/{author_id} [post]
func (h *Handler) newUserAuthor(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	authorID, err := strconv.Atoi(c.Param("author_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: id")
		return
	}

	if err := h.services.NewUserAuthor(userID, authorID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

//	@Summary		updateUser
//	@Description	Update User by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int					true	"user id"
//	@Param			input	body	models.UserUpdate	true	"user update model"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/{user_id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	var input models.UserUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateUser(id, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

//	@Summary		removeUserBook
//	@Description	Remove Book from User list
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int	true	"user id"
//	@Param			book_id	path	int	true	"book id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/{user_id}/books/{book_id} [delete]
func (h *Handler) removeUserBook(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: user_id")
		return
	}

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: book_id")
		return
	}

	if err := h.services.RemoveUserBook(userID, bookID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})

}

//	@Summary		removeUserAuthor
//	@Description	Remove Author from User list
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user_id		path	int	true	"user id"
//	@Param			author_id	path	int	true	"author id"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	statusResponse	"status response"
//	@Failure		400	{object}	errorResponse
//	@Failure		404	{object}	errorResponse
//	@Failure		500	{object}	errorResponse
//	@Router			/api/users/{user_id}/authors/{author_id} [delete]
func (h *Handler) removeUserAuthor(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: user_id")
		return
	}

	authorID, err := strconv.Atoi(c.Param("author_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid key: author_id")
		return
	}

	if err := h.services.RemoveUserAuthor(userID, authorID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
