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
	Provider   string
	Credential string
	Model      string
	IsDefault  bool
}

type Config struct {
	llms map[string]*llm
}

func New() (*Config, error) {
	return &Config{
		llms: make(map[string]*llm),
	}, nil

}

func LoadOrNew() (cfg *Config, isNew bool, err error) {
	cfg, err = Load()
	if err != nil {
		if !errors.Is(err, ErrConfigNotFound) {
			return nil, false, err
		}
	}

	if cfg == nil {
		cfg, err = New()
		if err != nil {
			return nil, false, err
		}
		isNew = true
	}

	return cfg, isNew, nil
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
			if llm.IsDefault {
				llm.IsDefault = false
				hasDefault = true
				break
			}
		}
	}

	if !isDefault && !hasDefault {
		return fmt.Errorf("must have at least one default LLM")
	}

	c.llms[provider] = &llm{
		Provider:   provider,
		Credential: credential,
		Model:      model,
		IsDefault:  isDefault,
	}

	return nil
}

func (c *Config) LLM(provider string) (*llm, bool) {
	llm, ok := c.llms[provider]
	return llm, ok
}

func (c *Config) DefaultLLM() (*llm, bool) {
	for _, llm := range c.llms {
		if llm.IsDefault {
			return llm, true
		}
	}
	return nil, false
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
				Provider:   llm.Provider,
				Credential: llm.Credential,
				Model:      llm.Model,
				Active:     llm.IsDefault,
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
			Provider:   data.Provider,
			Credential: data.Credential,
			Model:      data.Model,
			IsDefault:  data.Active,
		}
	}

	return nil
}

func (c *Config) Save() error {
	var errs []error

	for _, llm := range c.llms {
		if llm.Provider == "" {
			errs = append(errs, errors.New("provider cannot be empty"))
		}

		if llm.Model == "" {
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
