package main

import (
	"books/command"
	"books/common/logger"
	"books/config"
	localDb "books/db"
	"context"

	"fmt"
	"os"

	"github.com/spf13/viper"
)

func readConfig() {
	var err error

	viper.SetConfigFile(".env")
	viper.SetConfigType("props")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file .env not found")
	} else {
		viper.SetConfigFile(".env")
		viper.SetConfigType("props")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Override config parameters from environment variables if specified
	err = viper.Unmarshal(&config.Config)
	for _, key := range viper.AllKeys() {
		viper.BindEnv(key)
	}
}

func checkDb() {
	db := localDb.NewEntDb()
	_, err := db.ExecContext(context.Background(), "select 1")
	if err != nil {
		panic(fmt.Errorf("error connecting to db: %w", err))
	}
}

func main() {
	readConfig()
	raventClient := logger.NewRavenClient()
	logger.NewLogger(raventClient)
	checkDb()
	command.Execute()
}
