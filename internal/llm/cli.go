package llm

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

// AskConfigsCli is a command line interface that asks the user for the configuration.
type AskConfigsCli interface {
	Execute() (Config, error)
}

type askConfigsCliImpl struct {
	provider Provider
}

func (a *askConfigsCliImpl) Execute() (Config, error) {
	questions := []*survey.Question{
		a.createApiKeyQuestion(),
		a.createModelQuestion(),
	}

	answers := make(map[string]interface{})
	if err := survey.Ask(questions, &answers); err != nil {
		return nil, err
	}

	apiKey, ok1 := answers["APIKey"].(string)
	model, ok2 := answers["Model"].(string)

	if ok1 || ok2 {
		return nil, errors.New("invalid input")
	}

	return a.provider.NewConfig(apiKey, model), nil
}

func (a *askConfigsCliImpl) createApiKeyQuestion() *survey.Question {
	prompt := survey.Password{
		Message: a.provider.GetApiKeyLabel(),
		Help:    a.provider.GetApiKeyHelpText(),
	}

	return &survey.Question{
		Name:   "APIKey",
		Prompt: &prompt,
		Validate: func(ans interface{}) error {
			apiKey, ok := ans.(string)
			if !ok {
				return errors.New("invalid API key")
			}
			return ValidateApiKey(apiKey, a.provider)
		},
	}
}

func (a *askConfigsCliImpl) createModelQuestion() *survey.Question {
	prompt := survey.Select{
		Message: "Model",
		Options: a.provider.GetAllowedModels(),
		Help:    a.provider.GetModelHelpText(),
		Default: a.provider.GetDefaultModel(),
	}

	return &survey.Question{
		Name:   "Model",
		Prompt: &prompt,
		Validate: func(ans interface{}) error {
			model, ok := ans.(string)
			if !ok {
				return errors.New("invalid model")
			}
			return a.provider.ValidateModel(model)
		},
	}
}

func NewAskConfigsCli(provider Provider) AskConfigsCli {
	return &askConfigsCliImpl{provider: provider}
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
