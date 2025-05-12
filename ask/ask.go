package ask

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sashabaranov/go-openai"
)

func Provider(defaultValue string) (string, error) {
	if defaultValue == "" {
		defaultValue = "openai"
	}

	var provider string
	if err := survey.AskOne(
		&survey.Select{
			Message: "Provider:",
			Options: []string{"openai", "anthropic", "google"},
			Default: defaultValue,
			Help:    "The provider of the LLM.",
			VimMode: true,
		},
		&provider,
		survey.WithValidator(survey.Required),
	); err != nil {
		return "", err
	}

	return provider, nil
}

func Credential(defaultValue string) (string, error) {
	var credential string
	if err := survey.AskOne(
		&survey.Input{
			Message: "Credential:",
			Help:    "The credential for the LLM provider, e.g., OpenAI API key.",
			Default: defaultValue,
		},
		&credential,
		survey.WithValidator(survey.Required),
	); err != nil {
		return "", err
	}

	return credential, nil
}

func Model(provider string, defaultValue string) (string, error) {
	var model string
	if err := survey.AskOne(
		&survey.Select{
			Message: "Model:",
			Options: models(provider),
			Default: defaultValue,
			Help:    "The model to use.",
			VimMode: true,
		},
		&model,
		survey.WithValidator(survey.Required),
	); err != nil {
		return "", err
	}

	return model, nil
}

func IsDefault(defaultValue bool) (bool, error) {
	var isDefault bool
	if err := survey.AskOne(
		&survey.Confirm{
			Message: "Is this the default provider?",
			Default: defaultValue,
			Help:    "Whether this provider should be used by default.",
		},
		&isDefault,
	); err != nil {
		return false, err
	}

	return isDefault, nil
}

func models(provider string) []string {
	switch provider {
	case "openai":
		return []string{
			openai.GPT4o,
			openai.GPT4Dot1,
			openai.GPT4Dot1Mini,
			openai.GPT4Dot1Nano,
			openai.O1,
			openai.O1Mini,
			openai.O3,
			openai.O3Mini,
			openai.O4Mini,
		}
	default:
		panic(fmt.Sprintf("unsupported provider: %s", provider))
	}
}
