package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/models"
	"github.com/sirupsen/logrus"
)

type getAllBooksResponse struct {
	Data []models.Book `json:"data"`
}

type getAllAuthorsResponse struct {
	Data []models.Author `json:"data"`
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
