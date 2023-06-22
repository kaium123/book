package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(migrationApplyCmd)
}

var migrationApplyCmd = &cobra.Command{
	Use:   "migrate-apply",
	Short: "Apply migration",
	Run: func(cmd *cobra.Command, args []string) {
		dbUrl := viper.GetString("DB_URL")
		if booksSchema := viper.GetString("books_SCHEMA"); booksSchema != "" {
			dbUrl += " search_path=" + booksSchema
		}
		fmt.Println(dbUrl)
		execute := exec.Command("atlas", "migrate", "apply", "--dir", "file://ent/migrate/migrations/", "--url", convertConnectionString(dbUrl))
		execute.Stdout = os.Stdout
		execute.Stderr = os.Stderr

		// Run still runs the command and waits for completion
		// but the output is instantly piped to Stdout
		if err := execute.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}

func convertConnectionString(connStr string) string {
	// Split the connection string by space separator
	parts := strings.Split(connStr, " ")

	// Declare variables to store each part of the connection string
	var host, user, password, dbname, sslmode, port, searchPath string

	// Loop over the parts and extract the values
	for _, part := range parts {
		switch {
		case strings.HasPrefix(part, "host="):
			host = strings.TrimPrefix(part, "host=")
		case strings.HasPrefix(part, "user="):
			user = strings.TrimPrefix(part, "user=")
		case strings.HasPrefix(part, "password="):
			password = strings.TrimPrefix(part, "password=")
		case strings.HasPrefix(part, "dbname="):
			dbname = strings.TrimPrefix(part, "dbname=")
		case strings.HasPrefix(part, "sslmode="):
			sslmode = strings.TrimPrefix(part, "sslmode=")
		case strings.HasPrefix(part, "port="):
			port = strings.TrimPrefix(part, "port=")
		case strings.HasPrefix(part, "search_path="):
			searchPath = strings.TrimPrefix(part, "search_path=")
		}
	}

	// Construct the new connection string
	connStrNew := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
	if searchPath != "" {
		connStrNew = fmt.Sprintf("%s&search_path=%s", connStrNew, searchPath)
	}

	return connStrNew
}
