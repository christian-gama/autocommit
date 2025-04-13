package llm

import (
	"github.com/AlecAivazis/survey/v2"
)

// AskConfigsCli is a command line interface that asks the user for the configuration.
type AskConfigsCli interface {
	Execute() (Config, error)
}

type AskToChangeModelCli interface {
	Execute() (bool, error)
}

type askToChangeModelCliImpl struct{}

func (a *askToChangeModelCliImpl) Execute() (bool, error) {
	var provider Provider
	questions := CreateQuestions(
		provider,
		a.createModelQuestion,
	)

	type Answers struct {
		ChangeModel bool
	}

	var answers Answers

	err := survey.Ask(questions, &answers)
	if err != nil {
		return false, err
	}

	return answers.ChangeModel, nil
}

func (a *askToChangeModelCliImpl) createModelQuestion(provider Provider) *survey.Question {
	prompt := survey.Confirm{
		Message: "You reached the maximum number of tokens, but there is a model that can generate longer messages. Do you want to temporarily change the model?",
		Help:    "A model have a limited amount of tokens that can be generated at once. If you want to generate longer messages, you can temporarily change the model.",
		Default: true,
	}

	return &survey.Question{
		Name:   "ChangeModel",
		Prompt: &prompt,
	}
}

func NewAskToChangeModelCli() AskToChangeModelCli {
	return &askToChangeModelCliImpl{}
}
