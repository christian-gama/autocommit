package openai

import (
	"errors"
)

// ChatCommand is the interface that wraps the basic Execute method.
type ChatCommand interface {
	// Execute returns the response from the AI.
	Execute(config *Config, system *System, input string) (string, error)
}

// chatCommandImpl is an implementation of ChatCommand.
type chatCommandImpl struct {
	chat Chat
}

// Execute implements the ChatCommand interface.
func (c *chatCommandImpl) Execute(config *Config, system *System, input string) (string, error) {
	response, err := c.chat.Response(config, system, input)
	if err != nil {
		return "", err
	}
	return response, nil
}

// NewChatCommand creates a new instance of ChatCommand.
func NewChatCommand(chat Chat) ChatCommand {
	return &chatCommandImpl{
		chat: chat,
	}
}

// VerifyConfigCommand is the interface that wraps the basic Execute method.
type VerifyConfigCommand interface {
	// Execute will verify if the configs were initialized and if not, it will initialize them.
	Execute(getConfigsFn func() (*Config, error)) (*Config, error)
}

// verifyConfigCommandImpl is an implementation of VerifyConfigCommand.
type verifyConfigCommandImpl struct {
	repo ConfigRepo
}

// Execute Implements the VerifyConfigCommand interface.
func (v *verifyConfigCommandImpl) Execute(
	getConfigsFn func() (*Config, error),
) (config *Config, err error) {
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
func NewVerifyConfigCommand(repo ConfigRepo) VerifyConfigCommand {
	return &verifyConfigCommandImpl{
		repo: repo,
	}
}

// ResetConfigCommand is the interface that wraps the basic Execute method.
type ResetConfigCommand interface {
	// Execute will reset the configs.
	Execute() error
}

// resetConfigCommandImpl is an implementation of ResetConfigCommand.
type resetConfigCommandImpl struct {
	repo ConfigRepo
}

// Execute Implements the ResetConfigCommand interface.
func (r *resetConfigCommandImpl) Execute() error {
	if !r.repo.Exists() {
		return nil
	}

	return r.repo.DeleteConfig()
}

// NewResetConfigCommand creates a new instance of ResetConfigCommand.
func NewResetConfigCommand(repo ConfigRepo) ResetConfigCommand {
	return &resetConfigCommandImpl{
		repo: repo,
	}
}

// UpdateConfigCommand is the interface that wraps the basic Execute method.
type UpdateConfigCommand interface {
	// Execute will update the configs.
	Execute(config *Config) error
}

// updateConfigCommandImpl is an implementation of UpdateConfigCommand.
type updateConfigCommandImpl struct {
	repo ConfigRepo
}

// Execute Implements the UpdateConfigCommand interface.
func (u *updateConfigCommandImpl) Execute(config *Config) error {
	savedConfig, err := u.repo.GetConfig()
	if err != nil {
		return err
	}

	if savedConfig == nil {
		return errors.New("Configs weren't initialized yet - skipping...")
	}

	if config.ApiKey != "" {
		if err := ValidateApiKey(config.ApiKey); err != nil {
			return err
		}

		savedConfig.ApiKey = config.ApiKey
	}

	if config.Model != "" {
		if err := ValidateModel(config.Model); err != nil {
			return err
		}

		savedConfig.Model = config.Model
	}

	if config.Temperature != 0 {
		if err := ValidateTemperature(config.Temperature); err != nil {
			return err
		}

		savedConfig.Temperature = config.Temperature
	}

	return u.repo.UpdateConfig(savedConfig)
}

// NewUpdateConfigCommand creates a new instance of UpdateConfigCommand.
func NewUpdateConfigCommand(repo ConfigRepo) UpdateConfigCommand {
	return &updateConfigCommandImpl{
		repo: repo,
	}
}
