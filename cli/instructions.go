// Package cli provides the command-line interface for the autocommit tool.
package cli

import (
	"github.com/christian-gama/autocommit/instruction"
	"github.com/spf13/cobra"
)

// restoreInstruction is a command that resets the instructions file
// back to its default content.
var restoreInstruction = &cobra.Command{
	Use:                   "restore",
	Short:                 "Restore the instructions file to its default state",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cobra.NoFileCompletions,
	Run: func(cmd *cobra.Command, args []string) {
		if err := instruction.Restore(); err != nil {
			cmd.PrintErrf("Error restoring instructions file: %v\n", err)
			return
		}
		cmd.Println("✅ Instructions file has been restored to default")
	},
}

// Instruction is a command that opens the instructions file in the default text editor,
// allowing users to view and modify the template used for generating commit messages.
var Instruction = &cobra.Command{
	Use:                   "instructions",
	Short:                 "Open the instructions file",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if err := instruction.Open(); err != nil {
			cmd.PrintErrf("Error opening instructions file: %v\n", err)
			return
		}
	},
	ValidArgsFunction: cobra.NoFileCompletions,
}
