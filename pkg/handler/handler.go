package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		books := api.Group("/books")
		{
			books.GET("/", h.getAllBooks)
			books.POST("/", h.addBook)
			books.GET("/:id", h.getBook)
			books.PUT("/:id", h.updateBook)
			books.DELETE("/:id", h.deleteBook)
		}
		authors := api.Group("/authors")
		{
			authors.GET("/", h.getAllAuthors)
			authors.POST("/", h.addAuthor)
			authors.GET("/:id", h.getAuthor)
			authors.PUT("/:id", h.updateAuthor)
			authors.DELETE("/:id", h.deleteAuthor)
		}
	}

	return router
}
