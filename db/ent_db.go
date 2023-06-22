package db

import (
	"books/common/logger"
	"books/ent"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func NewEntDb() *ent.Client {
	dbUrl := viper.GetString("DB_URL")
	if booksSchema := viper.GetString("books_SCHEMA"); booksSchema != "" {
		dbUrl += " search_path=" + booksSchema
	}
	fmt.Println(dbUrl)
	drv, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	options := make([]ent.Option, 0)
	options = append(options, ent.Driver(drv))
	options = append(options, ent.Log(func(a ...any) {
		logger.LogDebug(a)
	}))

	if viper.GetString("GIN_MODE") == "debug" {
		options = append(options, ent.Debug())
	}
	entClient := ent.NewClient(options...)
	return entClient
}
