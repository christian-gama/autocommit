package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration file",
	Run:   runReset,
}

func runReset(cmd *cobra.Command, args []string) {
	if err := resetConfigCommand.Execute(); err != nil {
		panic(err)
	}

	fmt.Println("✅ Configuration file reset successfully - Next time you run autocommit, you will be asked to configure it again.")
}
