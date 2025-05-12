package cli

import (
	"github.com/christian-gama/autocommit/instruction"
	"github.com/spf13/cobra"
)

var restoreInstruction = &cobra.Command{
	Use:   "restore",
	Short: "Restore the instructions file to its default state",
	Run: func(cmd *cobra.Command, args []string) {
		if err := instruction.Restore(); err != nil {
			cmd.PrintErrf("Error restoring instructions file: %v\n", err)
			return
		}
		cmd.Println("Instructions file has been restored to default")
	},
}

var Instruction = &cobra.Command{
	Use:   "instructions",
	Short: "Open the instructions file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := instruction.Open(); err != nil {
			cmd.PrintErrf("Error opening instructions file: %v\n", err)
			return
		}
	},
}
