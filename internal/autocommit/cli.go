package autocommit

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/christian-gama/autocommit/internal/helpers"
)

// PostCommitCli is an interface for executing post commit CLI operations.
type PostCommitCli interface {
	// Execute executes the post commit CLI.
	Execute() (string, error)
}

// postCommitCliImpl is an implementation of PostCommitCli.
type postCommitCliImpl struct{}

// Execute implements the PostCommitCli interface.
func (p *postCommitCliImpl) Execute() (string, error) {
	questions := helpers.CreateQuestions(p.createActionQuestion)

	var option string

	err := survey.Ask(questions, &option)
	if err != nil {
		return "", err
	}

	return option, nil
}

func (p *postCommitCliImpl) createActionQuestion() *survey.Question {
	prompt := survey.Select{
		Message: "What would you like to do?",
		Help:    "Pick an option that you would like to do",
		Options: []string{
			CommitChangesOption,
			RegenerateOption,
			CopyToClipboardOption,
			AddInstructionOption,
			ExitOption,
		},
	}

	return &survey.Question{Name: "Option", Prompt: &prompt, Validate: survey.Required}
}

// NewPostCommitCli creates a new instance of PostCommitCli.
func NewPostCommitCli() PostCommitCli {
	return &postCommitCliImpl{}
}

const (
	CommitChangesOption   = "Commit changes to git"
	RegenerateOption      = "Regenerate commit message"
	CopyToClipboardOption = "Copy message to clipboard"
	AddInstructionOption  = "Add instruction"
	ExitOption            = "Exit"
)

// AddInstructionCli is an interface for executing add instruction CLI operations.
type AddInstructionCli interface {
	// Execute executes the add instruction CLI.
	Execute() (string, error)
}

// addInstructionCliImpl is an implementation of AddInstructionCli.
type addInstructionCliImpl struct{}

// Execute implements the AddInstructionCli interface.
func (a *addInstructionCliImpl) Execute() (string, error) {
	questions := helpers.CreateQuestions(a.createActionQuestion)

	var option string

	err := survey.Ask(questions, &option)
	if err != nil {
		return "", err
	}

	return option, nil
}

func (a *addInstructionCliImpl) createActionQuestion() *survey.Question {
	prompt := survey.Input{
		Message: "Add your your instruction",
		Help:    "This instruction will be used to regenerate your commit message based on your instructions.",
	}

	return &survey.Question{Name: "Option", Prompt: &prompt, Validate: survey.Required}
}

// NewAddInstructionCli creates a new instance of AddInstructionCli.
func NewAddInstructionCli() AddInstructionCli {
	return &addInstructionCliImpl{}
}
