package route

import (
	"books/book/controller"
	"books/book/repository"
	"books/book/service"
	"books/common/logger"
	"books/db"

	"github.com/gin-gonic/gin"
)

func BookSetup(api *gin.RouterGroup) {
	db := db.NewEntDb()
	raventClient := logger.NewRavenClient()
	logger := logger.NewLogger(raventClient)
	repo := repository.NewBookRepository(db, logger)
	service := service.NewBookService(repo)
	BookController := controller.NewBookController(service)

	Book := api.Group("/book")

	Book.GET("/:book_id", BookController.GetBook)
	Book.POST("/create", BookController.CreateBook)
	Book.POST("/update/:book_id", BookController.UpdateBook)
	Book.GET("/list", BookController.AllBook)
	Book.DELETE("/delete/:book_id", BookController.DeleteBook)
}
