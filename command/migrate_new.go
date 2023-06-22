package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrationNewCmd)
}

var migrationNewCmd = &cobra.Command{
	Use:   "migrate-new",
	Short: "Apply migration",
	Run: func(cmd *cobra.Command, args []string) {
		execute := exec.Command("atlas", "migrate", "new", "--dir", "file://ent/migrate/migrations/", args[0])
		execute.Stdout = os.Stdout
		execute.Stderr = os.Stderr

		// Run still runs the command and waits for completion
		// but the output is instantly piped to Stdout
		if err := execute.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}
