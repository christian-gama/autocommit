// Package cli provides the command-line interface for the autocommit tool.
package cli

import (
	"github.com/christian-gama/autocommit/v2/instruction"
	"github.com/spf13/cobra"
)

// restoreInstructionsCmd is a command that resets the instructions file
// back to its default content.
var restoreInstructionsCmd = &cobra.Command{
	Use:                   "restore",
	Short:                 "Restore the instructions file to its default state",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRestoreInstructions()
	},
}

func runRestoreInstructions() error {
	return instruction.Restore()
}

// instructionsCmd is a command that opens the instructions file in the default text editor,
// allowing users to view and modify the template used for generating commit messages.
var instructionsCmd = &cobra.Command{
	Use:                   "instructions",
	Short:                 "Open the instructions file",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInstructions()
	},
}

func runInstructions() error {
	return instruction.Open()
}
