package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

const (
	_dirName        = ".autocommit"
	_configFileName = "config.json"
)

var ErrConfigNotFound = errors.New("config file not found")

type llm struct {
	provider   string
	credential string
	model      string
	isDefault  bool
}

func (l *llm) Provider() string {
	return l.provider
}

func (l *llm) Credential() string {
	return l.credential
}

func (l *llm) Model() string {
	return l.model
}

func (l *llm) IsDefault() bool {
	return l.isDefault
}

type Config struct {
	llms map[string]*llm
}

func New() (*Config, error) {
	return &Config{
		llms: make(map[string]*llm),
	}, nil

}

func Load() (*Config, error) {
	config, err := New()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(config.configPath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrConfigNotFound
		}

		return nil, err
	}

	if err := config.Unmarshal(content); err != nil {
		return nil, errors.New("invalid config file")
	}

	return config, nil
}

func (c *Config) SetLLM(provider, model, credential string, isDefault bool) error {
	var hasDefault bool

	if isDefault {
		for _, llm := range c.llms {
			if llm.isDefault {
				llm.isDefault = false
				hasDefault = true
				break
			}
		}
	}

	if !isDefault && !hasDefault {
		return fmt.Errorf("must have at least one default LLM")
	}

	c.llms[provider] = &llm{
		provider:   provider,
		credential: credential,
		model:      model,
		isDefault:  isDefault,
	}

	return nil
}

func (c *Config) LLM(provider string) (*llm, error) {
	llm, ok := c.llms[provider]
	if !ok {
		return nil, fmt.Errorf("llm %s not found", provider)
	}

	return llm, nil
}

func (c *Config) DefaultLLM() (*llm, error) {
	for _, llm := range c.llms {
		if llm.isDefault {
			return llm, nil
		}
	}

	return nil, errors.New("no active LLM found")
}

func (c *Config) Marshal() ([]byte, error) {
	type llmData struct {
		Provider   string `json:"provider"`
		Credential string `json:"credential"`
		Model      string `json:"model"`
		Active     bool   `json:"active"`
	}

	var configData struct {
		LLMS []*llmData `json:"llms"`
	}

	for _, llm := range c.llms {
		configData.LLMS = append(
			configData.LLMS, &llmData{
				Provider:   llm.provider,
				Credential: llm.credential,
				Model:      llm.model,
				Active:     llm.isDefault,
			},
		)
	}

	data, err := json.Marshal(configData)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	return data, nil
}

func (c *Config) Unmarshal(data []byte) error {
	type llmData struct {
		Provider   string `json:"provider"`
		Credential string `json:"credential"`
		Model      string `json:"model"`
		Active     bool   `json:"active"`
	}

	var configData struct {
		LLMS []*llmData `json:"llms"`
	}

	if err := json.Unmarshal(data, &configData); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	for _, data := range configData.LLMS {
		c.llms[data.Provider] = &llm{
			provider:   data.Provider,
			credential: data.Credential,
			model:      data.Model,
			isDefault:  data.Active,
		}
	}

	return nil
}

func (c *Config) Save() error {
	var errs []error

	for _, llm := range c.llms {
		if llm.provider == "" {
			errs = append(errs, errors.New("provider cannot be empty"))
		}

		if llm.model == "" {
			errs = append(errs, errors.New("model cannot be empty"))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	if _, err := os.Stat(Dir()); os.IsNotExist(err) {
		return os.MkdirAll(Dir(), os.ModePerm)
	}

	if _, err := os.Stat(c.configPath()); err == nil {
		if err := os.Remove(c.configPath()); err != nil {
			return err
		}
	}

	content, err := c.Marshal()
	if err != nil {
		return err
	}

	return os.WriteFile(c.configPath(), content, os.ModePerm)
}

func (c *Config) Remove() error {
	c.llms = make(map[string]*llm)

	if _, err := os.Stat(c.configPath()); os.IsNotExist(err) {
		return nil
	}

	if err := os.Remove(c.configPath()); err != nil {
		return err
	}

	return nil
}

func (c *Config) configPath() string {
	return path.Join(Dir(), _configFileName)
}

func Dir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return path.Join(home, _dirName)
}
