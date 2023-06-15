package openai

import (
	"errors"
)

type ChatCommand interface {
	Execute(system *System, input string) (string, error)
}

type chatCommandImpl struct {
	chat Chat
}

func (c *chatCommandImpl) Execute(system *System, input string) (string, error) {
	response, err := c.chat.Response(system, input)
	if err != nil {
		return "", err
	}
	return response, nil
}

func NewChatCommand(chat Chat) ChatCommand {
	return &chatCommandImpl{
		chat: chat,
	}
}

type VerifyConfigCommand interface {
	Execute(getConfigsFn func() (*Config, error)) error
}

type verifyConfigCommandImpl struct {
	repo Repo
}

func (v *verifyConfigCommandImpl) Execute(getConfigsFn func() (*Config, error)) error {
	ok := v.repo.Exists()
	if !ok {
		configsInput, err := getConfigsFn()
		if err != nil {
			return err
		}

		if err := v.repo.SaveConfig(configsInput); err != nil {
			return err
		}
	}

	return nil
}

func NewVerifyConfigCommand(repo Repo) VerifyConfigCommand {
	return &verifyConfigCommandImpl{
		repo: repo,
	}
}

type ResetConfigCommand interface {
	Execute() error
}

type resetConfigCommandImpl struct {
	repo Repo
}

func (r *resetConfigCommandImpl) Execute() error {
	if !r.repo.Exists() {
		return nil
	}

	return r.repo.DeleteConfig()
}

func NewResetConfigCommand(repo Repo) ResetConfigCommand {
	return &resetConfigCommandImpl{
		repo: repo,
	}
}

type UpdateConfigCommand interface {
	Execute(config *Config) error
}

type updateConfigCommandImpl struct {
	repo Repo
}

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

	return u.repo.UpdateConfig(config)
}

func NewUpdateConfigCommand(repo Repo) UpdateConfigCommand {
	return &updateConfigCommandImpl{
		repo: repo,
	}
}
