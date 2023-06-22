package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrationHashCmd)
}

var migrationHashCmd = &cobra.Command{
	Use:   "migrate-hash",
	Short: "Hash migration",
	Run: func(cmd *cobra.Command, args []string) {
		execute := exec.Command("atlas", "migrate", "hash",
			"--dir", "file://ent/migrate/migrations/")
		execute.Stdout = os.Stdout
		execute.Stderr = os.Stderr

		// Run still runs the command and waits for completion
		// but the output is instantly piped to Stdout
		if err := execute.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}
