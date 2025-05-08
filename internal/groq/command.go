package groq

import (
	"errors"

	"github.com/christian-gama/autocommit/internal/llm"
)

type chatCommandImpl struct {
	chat llm.Chat
}

// Execute implements the ChatCommand interface.
func (c *chatCommandImpl) Execute(
	config llm.Config,
	system *llm.System,
	input string,
) (string, error) {
	openaiConfig, ok := config.(*GroqConfig)
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

type groqResetConfigCommandImpl struct {
	repo llm.ConfigRepo
}

func (c *groqResetConfigCommandImpl) Execute() error {
	if !c.repo.Exists() {
		return nil
	}

	return c.repo.DeleteConfig()
}

func NewGroqResetConfigCommand(repo llm.ConfigRepo) llm.ResetConfigCommand {
	return &groqResetConfigCommandImpl{
		repo: repo,
	}
}

// groqUpdateConfigCommandImpl is an implementation of UpdateConfigCommand.
type groqUpdateConfigCommandImpl struct {
	repo llm.ConfigRepo
}

// Execute Implements the UpdateConfigCommand interface.
func (u *groqUpdateConfigCommandImpl) Execute(config llm.Config) error {
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
func NewGroqUpdateConfigCommand(repo llm.ConfigRepo) llm.UpdateConfigCommand {
	return &groqUpdateConfigCommandImpl{
		repo: repo,
	}
}
