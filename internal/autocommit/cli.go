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
			CopyToClipboardOption,
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
	CopyToClipboardOption = fmt.Sprintf("%-2s Copy to clipboard", "📋")
	ExitOption            = fmt.Sprintf("%-2s Exit", "🚪")
)
