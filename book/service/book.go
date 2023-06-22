package service

import (
	"books/book/models"
	"books/book/repository"
	"books/common/logger"
	"books/common/utils"
	"context"
)

type BookServiceInterface interface {
	GetBook(id int, ctx context.Context) (*models.RespBook, error)
	CreateBook(Book *models.Book, ctx context.Context) (uint, error)
	AllBook(ctx context.Context) ([]models.RespBook, error)
	UpdateBook(id int, Book models.Book, ctx context.Context) (uint, error)
	DeleteBook(id int, ctx context.Context) (uint, error)
}

type BookService struct {
	repository repository.BookRepositoryInterface
}

func NewBookService(repository repository.BookRepositoryInterface) BookServiceInterface {
	return &BookService{repository: repository}
}

func (u BookService) GetBook(id int, ctx context.Context) (*models.RespBook, error) {

	resp, err := u.repository.FindByID(id, ctx)
	if err != nil {
		return nil, err
	}
	Book := models.RespBook{}
	utils.CopyStructToStruct(resp, &Book)
	return &Book, nil
}

func (u BookService) CreateBook(book *models.Book, ctx context.Context) (uint, error) {
	validationErr := book.Validate()
	if validationErr != nil {
		logger.LogError(validationErr)
		return 0, validationErr
	}

	id, err := u.repository.CreateBook(book, ctx)
	if err != nil {
		return 0, err
	}
	return id, err

}

func (u BookService) AllBook(ctx context.Context) ([]models.RespBook, error) {

	resp, err := u.repository.AllBook(ctx)
	if err != nil {
		return nil, err
	}

	Books := []models.RespBook{}
	for _, Book := range resp {
		tmpBook := &models.RespBook{}
		utils.CopyStructToStruct(Book, &tmpBook)
		Books = append(Books, *tmpBook)
	}
	return Books, nil

}

func (u BookService) UpdateBook(id int, book models.Book, ctx context.Context) (uint, error) {
	validationErr := book.Validate()
	if validationErr != nil {
		logger.LogError(validationErr)
		return 0, validationErr
	}

	_, err := u.repository.UpdateBook(id, book, ctx)
	if err != nil {
		return 0, err
	}
	return uint(id), nil

}

func (u BookService) DeleteBook(id int, ctx context.Context) (uint, error) {

	_, err := u.repository.DeleteBook(id, ctx)
	if err != nil {
		return 0, err
	}
	return uint(id), nil

}
