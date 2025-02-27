package openai

import (
	"errors"

	"github.com/christian-gama/autocommit/internal/llm"
)

// updateConfigCommandImpl is an implementation of UpdateConfigCommand.
type updateConfigCommandImpl struct {
	repo ConfigRepo
}

// Execute Implements the UpdateConfigCommand interface.
func (u *updateConfigCommandImpl) Execute(config llm.Config) error {
	openAIConfig, ok := config.(*Config)
	if !ok {
		return errors.New("invalid config type: expected OpenAI config")
	}

	savedConfig, err := u.repo.GetConfig()
	if err != nil {
		return err
	}

	openAISavedConfig, ok := savedConfig.(*Config)
	if !ok {
		return errors.New("invalid saved config type: expected OpenAI config")
	}

	if openAISavedConfig == nil {
		return errors.New("Configs weren't initialized yet - skipping...")
	}

	if openAIConfig.ApiKey != "" {
		if err := ValidateApiKey(openAIConfig.ApiKey); err != nil {
			return err
		}

		openAISavedConfig.ApiKey = openAIConfig.ApiKey
	}

	if openAIConfig.Model != "" {
		if err := ValidateModel(openAIConfig.Model); err != nil {
			return err
		}

		openAISavedConfig.Model = openAIConfig.Model
	}

	if openAIConfig.Temperature != 0 {
		if err := ValidateTemperature(openAIConfig.Temperature); err != nil {
			return err
		}

		openAISavedConfig.Temperature = openAIConfig.Temperature
	}

	return u.repo.UpdateConfig(openAISavedConfig)
}

// newUpdateConfigCommand creates a new instance of UpdateConfigCommand.
func newUpdateConfigCommand(repo llm.ConfigRepo) llm.UpdateConfigCommand {
	return &updateConfigCommandImpl{
		repo: repo.(ConfigRepo),
	}
}
