package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/christian-gama/autocommit/chat"
	"github.com/sashabaranov/go-openai"
)

type configAnswers struct {
	Model        string
	OpenAIAPIKey string
	Temperature  float32
}

func askUserForConfig() *configAnswers {
	openaiApiKeyPrompt := survey.Password{
		Message: "OpenAI API Key",
		Help:    "The API key to use for OpenAI. Get one at https://platform.openai.com/account/api-keys",
	}

	modelPrompt := createModelPrompt()

	temperaturePrompt := survey.Input{
		Message: "Temperature",
		Help:    "What sampling temperature to use",
		Default: "0.1",
	}

	var answers configAnswers

	err := survey.Ask([]*survey.Question{
		{Name: "OpenAIAPIKey", Prompt: &openaiApiKeyPrompt, Validate: survey.Required},
		{Name: "Model", Prompt: &modelPrompt, Validate: survey.Required},
		{
			Name:      "Temperature",
			Prompt:    &temperaturePrompt,
			Transform: convertToFloat32,
			Validate:  validateTemperature,
		},
	}, &answers)
	if err != nil {
		log.Fatalf("Failed to get config input: %v", err)
	}

	return &answers
}

func createModelPrompt() survey.Select {
	return survey.Select{
		Message: "Model name",
		Help:    "The model to use for completion",
		Default: openai.GPT3Dot5Turbo,
		Options: chat.Models,
	}
}

func convertToFloat32(ans interface{}) interface{} {
	value, err := strconv.ParseFloat(ans.(string), 32)
	if err != nil {
		return float32(0.28)
	}
	return float32(value)
}

func validateTemperature(ans interface{}) error {
	value, err := strconv.ParseFloat(ans.(string), 32)
	if err != nil {
		return fmt.Errorf("temperature must be a number")
	}
	if value <= 0 || value > 1 {
		return fmt.Errorf("temperature must be greater than 0 and less than or equal to 1")
	}
	return nil
}
