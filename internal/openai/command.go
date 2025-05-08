package openai

import (
	"errors"

	"github.com/christian-gama/autocommit/internal/llm"
)

// ChatCommand is the interface that wraps the basic Execute method.
// type ChatCommand interface {
// 	// Execute returns the response from the AI.
// 	Execute(config llm.Config, system *llm.System, input string) (string, error)
// }

// chatCommandImpl is an implementation of ChatCommand.
type chatCommandImpl struct {
	chat llm.Chat
}

// Execute implements the ChatCommand interface.
func (c *chatCommandImpl) Execute(
	config llm.Config,
	system *llm.System,
	input string,
) (string, error) {
	openaiConfig, ok := config.(*OpenAIConfig)
	if !ok {
		return "", errors.New("invallid config type")
	}
	response, err := c.chat.Response(openaiConfig, system, input)
	if err != nil {
		return "", err
	}
	return response, nil
}

// NewChatCommand creates a new instance of ChatCommand.
func NewChatCommand(chat llm.Chat) llm.ChatCommand {
	return &chatCommandImpl{
		chat: chat,
	}
}

// openaiResetConfigCommandImpl is an implementation of ResetConfigCommand.
type openaiResetConfigCommandImpl struct {
	repo llm.ConfigRepo
}

// Execute Implements the ResetConfigCommand interface.
func (r *openaiResetConfigCommandImpl) Execute() error {
	if !r.repo.Exists() {
		return nil
	}

	return r.repo.DeleteConfig()
}

// NewResetConfigCommand creates a new instance of ResetConfigCommand.
func NewOpenAIResetConfigCommand(repo llm.ConfigRepo) llm.ResetConfigCommand {
	return &openaiResetConfigCommandImpl{
		repo: repo,
	}
}

// UpdateConfigCommand is the interface that wraps the basic Execute method.
// type UpdateConfigCommand interface {
// 	// Execute will update the configs.
// 	Execute(config *OpenAIConfig) error
// }

// openaiUpdateConfigCommandImpl is an implementation of UpdateConfigCommand.
type openaiUpdateConfigCommandImpl struct {
	repo llm.ConfigRepo
}

// Execute Implements the UpdateConfigCommand interface.
func (u *openaiUpdateConfigCommandImpl) Execute(config llm.Config) error {
	savedConfig, err := u.repo.GetConfig()
	if err != nil {
		return err
	}

	if savedConfig == nil {
		return errors.New("configs weren't initialized yet")
	}

	if config.IsAPIKeySet() {
		apiKey := config.GetAPIKey()
		if err := ValidateApiKey(apiKey); err != nil {
			return err
		}

		savedConfig.SetAPIKey(apiKey)
	}

	if config.IsModelSet() {
		model := config.GetModel()
		if err := ValidateModel(model); err != nil {
			return err
		}

		savedConfig.SetModel(model)
	}

	if config.IsTemperatureSet() {
		temperature := config.GetTemperature()
		if err := ValidateTemperature(temperature); err != nil {
			return err
		}

		savedConfig.SetTemperature(temperature)
	}

	return u.repo.UpdateConfig(savedConfig)
}

// // NewUpdateConfigCommand creates a new instance of UpdateConfigCommand.
func NewOpenAIUpdateConfigCommand(repo llm.ConfigRepo) llm.UpdateConfigCommand {
	return &openaiUpdateConfigCommandImpl{
		repo: repo,
	}
}
