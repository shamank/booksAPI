package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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

}

func (h *Handler) getUserAuthors(c *gin.Context) {

}

func (h *Handler) newUserBook(c *gin.Context) {

}

func (h *Handler) newUserAuthor(c *gin.Context) {

}

func (h *Handler) updateUser(c *gin.Context) {

}
