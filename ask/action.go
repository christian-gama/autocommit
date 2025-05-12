package ask

import (
	"github.com/AlecAivazis/survey/v2"
)

type ActionOption string

const (
	ActionCommit          ActionOption = "✅ Commit"
	ActionAddInstruction  ActionOption = "📝 Add instruction"
	ActionRegenerate      ActionOption = "🔄 Regenerate"
	ActionCopyToClipboard ActionOption = "📋 Copy to clipboard"
	ActionExit            ActionOption = "🚪 Exit"
)

type Action struct{}

func NewAction() *Action {
	return &Action{}
}

func (a *Action) Action() (ActionOption, error) {
	var action string

	if err := survey.AskOne(
		&survey.Select{
			Message: "What do you want to do?",
			Options: []string{
				string(ActionCommit),
				string(ActionAddInstruction),
				string(ActionRegenerate),
				string(ActionCopyToClipboard),
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
