package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/models"
	"net/http"
	"strconv"
)

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
