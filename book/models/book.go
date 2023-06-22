package models

import (
	"time"

	"books/common/logger"
	appErr "books/errors"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear string `json:"publication_year"`
}

type RespBook struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear string    `json:"publication_year"`
	CreateAt        time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (e *Book) Validate() error {
	return v.ValidateStruct(e,
		v.Field(&e.Title,
			v.Required.ErrorObject(appErr.ApplicationError{
				ErrorType:      appErr.RequiredField,
				TranslationKey: "fieldRequired",
				TranslationParams: map[string]interface{}{
					"min":   2,
					"max":   1000,
					"field": "title",
				}}),
			v.Length(2, 1000).ErrorObject(appErr.ApplicationError{
				ErrorType:      appErr.LengthValidation,
				TranslationKey: "rangeValidation",
				TranslationParams: map[string]interface{}{
					"min":   2,
					"max":   1000,
					"field": "title",
				}}),
		),
		v.Field(&e.Author,
			v.Required.ErrorObject(appErr.ApplicationError{
				ErrorType:      appErr.RequiredField,
				TranslationKey: "fieldRequired",
				TranslationParams: map[string]interface{}{
					"min":   2,
					"max":   1000,
					"field": "author",
				}}),
			v.Length(2, 1000).ErrorObject(appErr.ApplicationError{
				ErrorType:      appErr.LengthValidation,
				TranslationKey: "rangeValidation",
				TranslationParams: map[string]interface{}{
					"min":   2,
					"max":   1000,
					"field": "author",
				}}),
		),
		v.Field(&e.PublicationYear,
			v.Length(2, 1000).ErrorObject(appErr.ApplicationError{
				ErrorType:      appErr.LengthValidation,
				TranslationKey: "rangeValidation",
				TranslationParams: map[string]interface{}{
					"min":   2,
					"max":   1000,
					"field": "author",
				}}),
			v.By(func(value interface{}) error {
				date := e.PublicationYear
				logger.LogError(date)
				publicationYear, err := StringToDate(date)
				if err != nil {
					return appErr.ApplicationError{
						ErrorType:      appErr.InvalidDate,
						TranslationKey: "invalidDate",
						TranslationParams: map[string]interface{}{
							"min":   0,
							"max":   1000,
							"field": "publication_year",
						}}
				}
				if publicationYear.After(time.Now()) {
					return appErr.ApplicationError{
						ErrorType:      appErr.InvalidDate,
						TranslationKey: "overDate",
						TranslationParams: map[string]interface{}{
							"min":   0,
							"max":   1000,
							"field": "publication_year",
						}}
				}
				return nil
			}),
		),
	)
}

func StringToDate(date string) (time.Time, error) {
	formattedDate, err := time.Parse(time.RFC3339, date)
	return formattedDate, err
}
