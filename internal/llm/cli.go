package llm

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
)

// AskConfigsCli is a command line interface that asks the user for the configuration.
type AskConfigsCli interface {
	Execute() (Config, error)
}

// askConfigsCliImpl is the implementation of AskConfigsCli.
type askConfigsCliImpl struct {
	provider Provider
}

func (a *askConfigsCliImpl) Execute() (Config, error) {
	// Get API Key
	var apiKey string
	apiKeyPrompt := a.createApiKeyQuestion().Prompt
	if err := survey.AskOne(apiKeyPrompt, &apiKey); err != nil {
		return nil, err
	}

	// Get Model
	var model string
	modelPrompt := a.createModelQuestion().Prompt
	if err := survey.AskOne(modelPrompt, &model); err != nil {
		return nil, err
	}

	var temperature float32
	temperaturePrompt := a.createTemperatureQuestion().Prompt
	if err := survey.AskOne(temperaturePrompt, &temperature); err != nil {
		return nil, err
	}

	return a.provider.NewConfig(apiKey, model, temperature), nil
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
		Message: a.provider.GetModelLabel(),
		Options: a.provider.GetAllowedModels(),
		Help:    a.provider.GetModelHelpText(),
		Default: a.provider.GetDefaultModel(),
	}

	return &survey.Question{
		Name:   "Model",
		Prompt: &prompt,
		Validate: func(ans interface{}) error {
			optionAnswer, ok := ans.(core.OptionAnswer)
			if !ok {
				return fmt.Errorf("unexpected type: %T, model type: %T", ans, optionAnswer.Value)
			}
			// do a type check and not assertion on optionAnswer.Value
			model := optionAnswer.Value

			return a.provider.ValidateModel(model)
		},
	}
}

func (a *askConfigsCliImpl) createTemperatureQuestion() *survey.Question {
	prompt := survey.Input{
		Message: "Temperature",
		Help:    a.provider.GetTemperatureHelpText(),
		Default: "0.5",
	}

	return &survey.Question{
		Name:   "Temperature",
		Prompt: &prompt,
		Validate: func(ans interface{}) error {
			// Ensure the answer is a string
			answerStr, ok := ans.(string)
			if !ok {
				return errors.New("invalid input: temperature must be a string")
			}

			// Convert the string to a float64
			temperature, err := strconv.ParseFloat(answerStr, 32)
			if err != nil {
				return errors.New("invalid input: temperature must be a valid number")
			}

			// Validate the temperature as a float32
			return a.provider.ValidateTemperature(float32(temperature))
		},
	}
}

func NewAskConfigsCli(provider Provider) AskConfigsCli {
	return &askConfigsCliImpl{
		provider: provider,
	}
}

// AskToChangeModelCli is a command line interface that asks the user if they want to change the model.
type AskToChangeModelCli interface {
	Execute() (bool, error)
}

// askToChangeModelCliImpl is the implementation of AskToChangeModelCli.
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
