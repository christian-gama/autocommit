package cli

import (
	"github.com/spf13/cobra"
)

var (
	verbose bool
	cmd     = &cobra.Command{
		Use: "autocommit",
		Run: run,
	}
)

func init() {
	cmd.PersistentFlags().
		BoolVarP(&verbose, "verbose", "v", false, "create a verbose commit message")
}

// Execute executes the root command.
func Execute() error {
	return cmd.Execute()
}
