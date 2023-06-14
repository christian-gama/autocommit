package cmd

import (
	"github.com/christian-gama/autocommit/config"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		config.NewStore().DeleteConfigFile()
	},
}
