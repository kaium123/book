package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrationInitSchemaCmd)
}

var migrationInitSchemaCmd = &cobra.Command{
	Use:   "entity-init-schema",
	Short: "Init Entity Schema",
	Run: func(cmd *cobra.Command, args []string) {
		execute := exec.Command("go", "run", "-mod=mod",
			"entgo.io/ent/cmd/ent", "new", args[0])
		execute.Stdout = os.Stdout
		execute.Stderr = os.Stderr

		// Run still runs the command and waits for completion
		// but the output is instantly piped to Stdout
		if err := execute.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}
