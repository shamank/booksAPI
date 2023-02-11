package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/models"
	"github.com/sirupsen/logrus"
)

type getAllBooksResponse struct {
	Books []models.Book `json:"books"`
}

type getAllAuthorsResponse struct {
	Authors []models.Author `json:"authors"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, msg string) {
	logrus.Error(msg)

	c.AbortWithStatusJSON(statusCode, errorResponse{Message: msg})
}
