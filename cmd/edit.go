package cmd

import (
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the system message",
	Run:   runEdit,
}

func runEdit(cmd *cobra.Command, args []string) {
	if err := systemMsgHealthCheck.Execute(); err != nil {
		panic(err)
	}

	if err := openSystemMsgCommand.Execute(); err != nil {
		panic(err)
	}
}
