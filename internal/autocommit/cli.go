package autocommit

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/christian-gama/autocommit/internal/helpers"
)

type PostCommitCli interface {
	Execute() (string, error)
}

type postCommitCliImpl struct{}

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
			ExitOption,
		},
	}

	return &survey.Question{Name: "Option", Prompt: &prompt, Validate: survey.Required}
}

func NewPostCommitCli() PostCommitCli {
	return &postCommitCliImpl{}
}

const (
	CommitChangesOption   = "Commit changes to git"
	RegenerateOption      = "Regenerate commit message"
	CopyToClipboardOption = "Copy message to clipboard"
	ExitOption            = "Exit"
)
