package cli

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

type chatAnswers struct {
	Option string
}

const (
	commitChangesOption                = "Commit changes to git"
	generateCommitMessageOption        = "Generate a new commit message"
	copyCommitMessageToClipboardOption = "Copy the commit message to clipboard"
	exitOption                         = "Exit"
)

func askUserForChatOption() *chatAnswers {
	optionPrompt := survey.Select{
		Message: "What would you like to do?",
		Help:    "Pick an option that you would like to do",
		Options: []string{
			commitChangesOption,
			generateCommitMessageOption,
			copyCommitMessageToClipboardOption,
			exitOption,
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
