package command

import (
	"books/ent/migrate"
	"context"
	"fmt"
	"log"
	"os"

	atlasMigrate "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(migrationDiffCmd)
}

var migrationDiffCmd = &cobra.Command{
	Use:   "migrate-diff",
	Short: "Run migration diff",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		dbUrl := viper.GetString("DB_URL")
		if booksSchema := viper.GetString("books_SCHEMA"); booksSchema != "" {
			dbUrl += " search_path=" + booksSchema
		}
		dbUrl = convertConnectionString(dbUrl)
		fmt.Println(dbUrl)
		dir, err := atlasMigrate.NewLocalDir("ent/migrate/migrations")
		if err != nil {
			log.Fatalf("failed creating atlas migration directory: %v", err)
		}
		// Migrate diff options.
		opts := []schema.MigrateOption{
			schema.WithDir(dir),                          // provide migration directory
			schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
			schema.WithDialect(dialect.Postgres),         // Ent dialect to use
			schema.WithFormatter(atlasMigrate.DefaultFormatter),
			// schema.WithDropColumn(true),
			// schema.WithDropIndex(true),
			// schema.WithDiffHook(renameColumnHook),
		}
		if len(os.Args) != 3 {
			log.Fatalln("migration name is required. Use: 'go run main.go migrate-diff <name>'")
		}
		// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
		err = migrate.NamedDiff(ctx, dbUrl, os.Args[2], opts...)
		if err != nil {
			log.Fatalf("failed generating migration file: %v", err)
		}
	},
}

// func renameColumnHook(next schema.Differ) schema.Differ {
// 	return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
// 		changes, err := next.Diff(current, desired)
// 		if err != nil {
// 			return nil, err
// 		}
// 		for _, c := range changes {
// 			m, ok := c.(*atlas.ModifyTable)
// 			// Skip if the change is not a ModifyTable,
// 			// or if the table is not the "users" table.
// 			if !ok || m.T.Name != user.Table {
// 				continue
// 			}
// 			changes := atlas.Changes(m.Changes)
// 			switch i, j := changes.IndexDropColumn("email2"), changes.IndexAddColumn("email"); {
// 			case i != -1 && j != -1:
// 				// Append a new renaming change.
// 				changes = append(changes, &atlas.RenameColumn{
// 					From: changes[i].(*atlas.DropColumn).C,
// 					To:   changes[j].(*atlas.AddColumn).C,
// 				})
// 				// Remove the drop and add changes.
// 				changes.RemoveIndex(i, j)
// 				m.Changes = changes
// 			case i != -1 || j != -1:
// 				return nil, errors.New("old_name and new_name must be present or absent")
// 			}
// 		}
// 		return changes, nil
// 	})
// }
