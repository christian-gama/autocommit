package autocommit

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/christian-gama/autocommit/internal/llm"
)

// PostCommitCli is an interface for executing post commit CLI operations.
type PostCommitCli interface {
	// Execute executes the post commit CLI.
	Execute() (string, error)
}

var provider llm.Provider

// postCommitCliImpl is an implementation of PostCommitCli.
type postCommitCliImpl struct{}

// Execute implements the PostCommitCli interface.
func (p *postCommitCliImpl) Execute() (string, error) {
	questions := llm.CreateQuestions(provider, p.createActionQuestion)

	var option string

	err := survey.Ask(questions, &option)
	if err != nil {
		return "", err
	}

	return option, nil
}

func (p *postCommitCliImpl) createActionQuestion(provider llm.Provider) *survey.Question {
	prompt := survey.Select{
		Message: "What would you like to do?",
		Help:    "Pick an option that you would like to do",
		VimMode: true,
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

var (
	CommitChangesOption   = fmt.Sprintf("%-2s Commit changes", "💾")
	RegenerateOption      = fmt.Sprintf("%-2s Regenerate", "🔄")
	CopyToClipboardOption = fmt.Sprintf("%-2s Copy to clipboard", "📋")
	AddInstructionOption  = fmt.Sprintf("%-4s Add instruction", "✏️")
	ExitOption            = fmt.Sprintf("%-2s Exit", "🚪")
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
	questions := llm.CreateQuestions(provider, a.createActionQuestion)

	var option string

	err := survey.Ask(questions, &option)
	if err != nil {
		return "", err
	}

	return option, nil
}

func (a *addInstructionCliImpl) createActionQuestion(provider llm.Provider) *survey.Question {
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
