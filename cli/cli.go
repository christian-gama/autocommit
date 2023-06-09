package cli

import (
	"github.com/christian-gama/autocommit/store"
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

	cmd.AddCommand(reset)
}

var reset = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		store.NewStore().DeleteConfigFile()
	},
}

// Execute executes the root command.
func Execute() error {
	return cmd.Execute()
}
