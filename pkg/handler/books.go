package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllBooks(c *gin.Context) {
	id, _ := c.Get(userCtx)

	c.JSON(http.StatusOK, map[string]interface{}{
		userCtx: id,
	})
}

func (h *Handler) addBook(c *gin.Context) {

}

func (h *Handler) getBook(c *gin.Context) {

}

func (h *Handler) updateBook(c *gin.Context) {

}

func (h *Handler) deleteBook(c *gin.Context) {

}
