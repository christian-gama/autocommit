package cli

import (
	"github.com/christian-gama/autocommit/store"
	"github.com/spf13/cobra"
)

var reset = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		store.NewStore().DeleteConfigFile()
	},
}
