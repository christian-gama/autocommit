// Package ask provides interactive prompt utilities for the autocommit tool.
package ask

import (
	"github.com/AlecAivazis/survey/v2"
)

// ActionOption represents an available action for users to choose from
// when interacting with the autocommit tool.
type ActionOption string

// Available ActionOptions for users to select.
const (
	// ActionCommit commits the changes with the generated message.
	ActionCommit ActionOption = "✅ Commit"
	// ActionCopyToClipboard copies the commit command to the clipboard.
	ActionCopyToClipboard ActionOption = "📋 Copy to clipboard"
	// ActionRegenerate regenerates a new commit message.
	ActionRegenerate ActionOption = "🔄 Regenerate"
	// ActionAddInstruction adds a custom instruction to guide message generation.
	ActionAddInstruction ActionOption = "📝 Add instruction"
	// ActionExit exits the application without committing.
	ActionExit ActionOption = "🚪 Exit"
)

// Action provides methods for prompting user actions related to commit messages.
type Action struct{}

// NewAction creates a new Action instance for prompting users for actions.
func NewAction() *Action {
	return &Action{}
}

// Action prompts the user to select what action to take with the generated
// commit message, such as committing, adding an instruction, regenerating,
// copying to clipboard, or exiting.
func (a *Action) Action() (ActionOption, error) {
	var action string

	if err := survey.AskOne(
		&survey.Select{
			Message: "What now?",
			Options: []string{
				string(ActionCommit),
				string(ActionCopyToClipboard),
				string(ActionRegenerate),
				string(ActionAddInstruction),
				string(ActionExit),
			},
			Default: string(ActionCommit),
			Help:    "The action to perform.",
			VimMode: true,
		},
		&action,
		survey.WithValidator(survey.Required),
	); err != nil {
		return "", err
	}

	return ActionOption(action), nil
}

// Instruction prompts the user to enter a custom instruction that will be
// used to guide the generation of a commit message.
func (a *Action) Instruction() (string, error) {
	var instruction string
	if err := survey.AskOne(
		&survey.Input{
			Message: "Instruction:",
			Help:    "The instruction to add.",
		},
		&instruction,
		survey.WithValidator(survey.Required),
	); err != nil {
		return "", err
	}

	return instruction, nil
}
