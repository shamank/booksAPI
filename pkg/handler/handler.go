package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamank/booksAPI/pkg/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/shamank/booksAPI/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/:user_id", h.getUserById)
			users.GET("/:user_id/books", h.getUserBooks)
			users.GET("/:user_id/authors", h.getUserAuthors)

			users.POST("/:user_id/books/:book_id", h.checkRightsForEdit, h.newUserBook)

			users.PUT("/:user_id/books/:book_id", h.checkRightsForEdit, h.setBookRating)

			users.POST("/:user_id/authors/:author_id", h.checkRightsForEdit, h.newUserAuthor)

			users.DELETE("/:user_id/books/:book_id", h.checkRightsForEdit, h.removeUserBook)
			users.DELETE("/:user_id/authors/:author_id", h.checkRightsForEdit, h.removeUserAuthor)

			users.PUT("/:user_id", h.checkRightsForEdit, h.updateUser)

		}
		books := api.Group("/books")
		{
			books.GET("/", h.getAllBooks)
			books.POST("/", h.checkOnModerator, h.createBook)
			books.GET("/:id", h.getBook)
			books.PUT("/:id", h.checkOnModerator, h.updateBook)
			books.DELETE("/:id", h.checkOnModerator, h.deleteBook)
		}
		authors := api.Group("/authors")
		{
			authors.GET("/", h.getAllAuthors)
			authors.POST("/", h.checkOnModerator, h.addAuthor)
			authors.GET("/:id", h.getAuthor)
			authors.PUT("/:id", h.checkOnModerator, h.updateAuthor)
			authors.DELETE("/:id", h.checkOnModerator, h.deleteAuthor)
		}
	}

	return router
}
