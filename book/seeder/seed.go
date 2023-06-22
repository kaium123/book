package seeder

import (
	"books/ent"
	"context"
	"fmt"
)

func Seed(db *ent.Client) {

	books := make([]*ent.BookCreate, 0)
	for i := 0; i < 50; i++ {
		book := db.Book.Create()
		books = append(books, book)
	}
	err := db.Book.CreateBulk(books...).Exec(context.Background())
	if err != nil {
		panic(fmt.Errorf("failed generating statement: %w", err))
	}
}
