package ask

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/christian-gama/autocommit/llm"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Provider() (string, error) {
	var provider string
	if err := survey.AskOne(
		&survey.Select{
			Message: "Provider:",
			Options: []string{llm.OpenAI, llm.Ollama2, llm.Mistral, llm.GoogleAI},
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

func (c *Config) Credential(defaultValue string) (string, error) {
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

func (c *Config) Model(provider string, defaultValue string) (string, error) {
	var defaultModel any

	if defaultValue != "" {
		defaultModel = defaultValue
	}

	var model string
	if err := survey.AskOne(
		&survey.Select{
			Message: "Model:",
			Options: llm.Models(provider),
			Default: defaultModel,
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

func (c *Config) IsDefault(defaultValue bool) (bool, error) {
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
