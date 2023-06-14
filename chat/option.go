package chat

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
	"github.com/christian-gama/autocommit/git"
)

type chatAnswers struct {
	Option string
}

const (
	CommitChangesOption                = "Commit changes to git"
	GenerateCommitMessageOption        = "Generate a new commit message"
	CopyCommitMessageToClipboardOption = "Copy the commit message to clipboard"
	ExitOption                         = "Exit"
)

func AskUserForChatOption() *chatAnswers {
	optionPrompt := survey.Select{
		Message: "What would you like to do?",
		Help:    "Pick an option that you would like to do",
		Options: []string{
			CommitChangesOption,
			GenerateCommitMessageOption,
			CopyCommitMessageToClipboardOption,
			ExitOption,
		},
	}

	var answers chatAnswers

	err := survey.Ask([]*survey.Question{
		{Name: "Option", Prompt: &optionPrompt, Validate: survey.Required},
	}, &answers)
	if err != nil {
		log.Fatalf("Failed to get user input: %v", err)
	}

	return &answers
}

func HandleChatOption(restart func(), option string, commitMessage string) {
	switch option {
	case CommitChangesOption:
		git.Commit(commitMessage)

	case GenerateCommitMessageOption:
		restart()

	case CopyCommitMessageToClipboardOption:
		if err := clipboard.WriteAll(fmt.Sprintf("git commit -m \"%s\"", commitMessage)); err != nil {
			log.Fatalf("Failed to copy commit message to clipboard: %v", err)
		}
	}
}
