package repository

import (
	"books/book/models"
	"books/common/logger"
	"books/ent"
	"books/ent/book"
	"books/errors"
	"context"
	"net/http"
)

type BookRepositoryInterface interface {
	FindByID(id int, ctx context.Context) (*ent.Book, error)
	CreateBook(Book *models.Book, ctx context.Context) (uint, error)
	AllBook(ctx context.Context) ([]*ent.Book, error)
	UpdateBook(id int, Book models.Book, ctx context.Context) (uint, error)
	DeleteBook(id int, ctx context.Context) (uint, error)
}

type BookRepository struct {
	Db     *ent.Client
	logger logger.LoggerInterface
}

func NewBookRepository(db *ent.Client, logger logger.LoggerInterface) BookRepositoryInterface {
	return &BookRepository{Db: db, logger: logger}
}

func (r *BookRepository) FindByID(id int, ctx context.Context) (*ent.Book, error) {
	resp, err := r.Db.Book.Query().
		Where(book.ID(id), book.IsDeleted(false)).
		First(ctx)
	if err != nil {
		return nil, errors.ApplicationError{TranslationKey: "failedFindingBook", HttpCode: http.StatusInternalServerError}
	}

	return resp, nil
}

func (r *BookRepository) CreateBook(book *models.Book, ctx context.Context) (uint, error) {
	publicationYear, _ := models.StringToDate(book.PublicationYear)
	resp, err := r.Db.Book.Create().
		SetAuthor(book.Author).
		SetTitle(book.Title).
		SetPublicationYear(publicationYear).
		Save(ctx)
	if err != nil {
		return 0, errors.ApplicationError{TranslationKey: "failedCreateBook", HttpCode: http.StatusInternalServerError}
	}

	return uint(resp.ID), nil
}


func (r *BookRepository) AllBook(ctx context.Context) ([]*ent.Book, error) {

	resp, err := r.Db.Book.Query().
		Where(book.IsDeleted(false)).
		All(ctx)
	if err != nil {
		return nil, errors.ApplicationError{TranslationKey: "failedFindingBook", HttpCode: http.StatusInternalServerError}
	}

	return resp, nil

}

func (r *BookRepository) UpdateBook(id int, book models.Book, ctx context.Context) (uint, error) {
	publicationYear, _ := models.StringToDate(book.PublicationYear)
	_, err := r.Db.Book.UpdateOneID(id).
		SetAuthor(book.Author).
		SetTitle(book.Title).
		SetPublicationYear(publicationYear).
		Save(ctx)
	if err != nil {
		return 0, errors.ApplicationError{TranslationKey: "failedUpdateBook", HttpCode: http.StatusInternalServerError}
	}

	return uint(id), nil
}

func (r *BookRepository) DeleteBook(id int, ctx context.Context) (uint, error) {
	_,err:=r.FindByID(id,ctx)
	if err != nil {
		return 0, err
	}
	
	_, err = r.Db.Book.UpdateOneID(id).
		SetIsDeleted(true).
		Save(ctx)
	if err != nil {
		return 0, errors.ApplicationError{TranslationKey: "failedDeleteBook", HttpCode: http.StatusInternalServerError}
	}

	return uint(id), nil
}
