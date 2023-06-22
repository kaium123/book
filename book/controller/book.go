package controller

import (
	"books/book/models"
	"books/book/service"
	"books/common/logger"
	"books/common/utils"
	"books/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BookController struct {
	errors.GinError
	service service.BookServiceInterface
}

func NewBookController(service service.BookServiceInterface) *BookController {
	return &BookController{service: service}
}

func (c BookController) GetBook(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()

	Book_id := ginContext.Param("book_id")
	logger.LogError(Book_id)
	id, err := strconv.Atoi(Book_id)
	if err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.Trans("invalidRequestParams", nil)})
		return
	}

	Book, err := c.service.GetBook(id, ctx)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}
	ginContext.JSON(http.StatusCreated, gin.H{"message": "Book", "Book": Book})

}

func (c BookController) CreateBook(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	Book := &models.Book{}

	if err := ginContext.ShouldBindBodyWith(&Book, binding.JSON); err != nil {
		logger.LogError("JSON body binding error ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalidJSONBody"})
		return
	}

	BookID, err := c.service.CreateBook(Book, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": utils.Trans("bookCreated", nil), "id": BookID})
}

func (c BookController) AllBook(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()

	Books, err := c.service.AllBook(ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": "Books", "Books": Books})
}

func (c BookController) UpdateBook(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()
	Book := models.Book{}

	Book_id := ginContext.Param("book_id")
	logger.LogError(Book_id)
	id, err := strconv.Atoi(Book_id)
	if err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.Trans("invalidRequestParams", nil)})
		return
	}

	if err := ginContext.ShouldBindBodyWith(&Book, binding.JSON); err != nil {
		logger.LogError("JSON body binding error ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalidJSONBody"})
		return
	}

	BookID, err := c.service.UpdateBook(id, Book, ctx)

	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message": utils.Trans("bookUpdated", nil), "id": BookID})
}

func (c BookController) DeleteBook(ginContext *gin.Context) {
	ctx := ginContext.Request.Context()

	Book_id := ginContext.Param("book_id")
	id, err := strconv.Atoi(Book_id)
	if err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.Trans("invalidRequestParams", nil)})
		return
	}

	BookID, err := c.service.DeleteBook(id, ctx)
	if err != nil {
		ginContext.AbortWithStatusJSON(c.GetStatusCode(err), gin.H{"error": c.ErrorTraverse(err)})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"message":  utils.Trans("bookDeleted", nil), "id": BookID})
}
