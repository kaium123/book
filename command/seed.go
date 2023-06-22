package command

import (

	//	"books/ent"

	//"fmt"

	"books/db"
	bookSeeder "books/book/seeder"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run sync",
	Run: func(cmd *cobra.Command, args []string) {
		db := db.NewEntDb()
		bookSeeder.Seed(db)
	},
}
