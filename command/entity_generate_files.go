package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrationGenerateCmd)
}

var migrationGenerateCmd = &cobra.Command{
	Use:   "entity-generate-files",
	Short: "Generate Entity Files",
	Run: func(cmd *cobra.Command, args []string) {
		execute := exec.Command("go", "generate", "./ent")
		execute.Stdout = os.Stdout
		execute.Stderr = os.Stderr

		// Run still runs the command and waits for completion
		// but the output is instantly piped to Stdout
		if err := execute.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}
