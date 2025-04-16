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

// VerifyConfigCommand is the interface that wraps the basic Execute method.
// type VerifyConfigCommand interface {
// 	// Execute will verify if the configs were initialized and if not, it will initialize them.
// 	Execute(getConfigsFn func() (*OpenAIConfig, error)) (*OpenAIConfig, error)
// }

// verifyConfigCommandImpl is an implementation of VerifyConfigCommand.
type verifyConfigCommandImpl struct {
	repo llm.ConfigRepo
}

// Execute Implements the VerifyConfigCommand interface.
func (v *verifyConfigCommandImpl) Execute(
	getConfigsFn func() (llm.Config, error),
) (config llm.Config, err error) {
	ok := v.repo.Exists()
	if !ok {
		config, err = getConfigsFn()
		if err != nil {
			return nil, err
		}

		if err := v.repo.SaveConfig(config); err != nil {
			return nil, err
		}
	} else {
		config, err = v.repo.GetConfig()
		if err != nil {
			return nil, err
		}
	}

	return config, err
}

// NewVerifyConfigCommand creates a new instance of VerifyConfigCommand.
// func NewVerifyConfigCommand(repo llm.ConfigRepo) llm.VerifyConfigCommand {
// 	return &verifyConfigCommandImpl{
// 		repo: repo,
// 	}
// }

// ResetConfigCommand is the interface that wraps the basic Execute method.
// type ResetConfigCommand interface {
// 	// Execute will reset the configs.
// 	Execute() error
// }

// resetConfigCommandImpl is an implementation of ResetConfigCommand.
// type resetConfigCommandImpl struct {
// 	repo llm.ConfigRepo
// }

// Execute Implements the ResetConfigCommand interface.
// func (r *resetConfigCommandImpl) Execute() error {
// 	if !r.repo.Exists() {
// 		return nil
// 	}
//
// 	return r.repo.DeleteConfig()
// }

// NewResetConfigCommand creates a new instance of ResetConfigCommand.
// func NewResetConfigCommand(repo llm.ConfigRepo) ResetConfigCommand {
// 	return &resetConfigCommandImpl{
// 		repo: repo,
// 	}
// }

// UpdateConfigCommand is the interface that wraps the basic Execute method.
// type UpdateConfigCommand interface {
// 	// Execute will update the configs.
// 	Execute(config *OpenAIConfig) error
// }

// updateConfigCommandImpl is an implementation of UpdateConfigCommand.
// type updateConfigCommandImpl struct {
// 	repo llm.ConfigRepo
// }

// Execute Implements the UpdateConfigCommand interface.
// func (u *updateConfigCommandImpl) Execute(config llm.Config) error {
// 	savedConfig, err := u.repo.GetConfig()
// 	if err != nil {
// 		return err
// 	}
//
// 	if savedConfig == nil {
// 		return errors.New("Configs weren't initialized yet - skipping...")
// 	}
//
// 	apiKey := config.GetAPIKey()
// 	if apiKey != "" {
// 		if err := ValidateApiKey(apiKey); err != nil {
// 			return err
// 		}
//
// 		savedConfig.SetAPIKey(apiKey)
// 	}
//
// 	model := config.GetModel()
// 	if model != "" {
// 		if err := ValidateModel(model); err != nil {
// 			return err
// 		}
//
// 		savedConfig.SetModel(model)
// 	}

// 	temperature := config.GetTemperature()
// 	if temperature != 0 {
// 		if err := ValidateTemperature(temperature); err != nil {
// 			return err
// 		}
//
// 		savedConfig.SetTemperature(temperature)
// 	}
//
// 	return u.repo.UpdateConfig(savedConfig)
// }
//
// // NewUpdateConfigCommand creates a new instance of UpdateConfigCommand.
// func NewUpdateConfigCommand(repo llm.ConfigRepo) llm.UpdateConfigCommand {
// 	return &updateConfigCommandImpl{
// 		repo: repo,
// 	}
// }
