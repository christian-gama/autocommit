// Package config provides configuration management for the autocommit tool.
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

// ErrConfigNotFound is returned when the configuration file cannot be found.
var ErrConfigNotFound = errors.New("config file not found")

type llmSettings struct {
	Provider   string `json:"provider"`
	Credential string `json:"credential"`
	Model      string `json:"model"`
	IsDefault  bool   `json:"is_default"`
}

// Config manages the configuration for LLM providers and their credentials.
type Config struct {
	llms map[string]*llmSettings
}

// New creates a new empty configuration instance.
func New() (*Config, error) {
	return &Config{
		llms: make(map[string]*llmSettings),
	}, nil

}

func HasConfig() bool {
	_, err := os.Stat(Dir())
	if os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(path.Join(Dir(), _configFileName))
	return !os.IsNotExist(err)
}

// LoadOrNew attempts to load an existing configuration file, or creates a new one
// if no configuration exists. Returns the config, whether it's new, and any error.
func LoadOrNew() (cfg *Config, err error) {
	cfg, err = Load()
	if err != nil {
		if !errors.Is(err, ErrConfigNotFound) {
			return nil, err
		}
	}

	if cfg == nil {
		cfg, err = New()
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

// Load reads the configuration file from disk and returns a populated Config.
// Returns ErrConfigNotFound if the configuration file doesn't exist.
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

// HasAnyLLM returns true if there are any LLM providers configured.
func (c *Config) HasAnyLLM() bool {
	return len(c.llms) > 0
}

// SetLLM configures an LLM provider with the given details.
// If isDefault is true, it will make this the default provider and unset any
// previous default. Ensures there is always at least one default provider.
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

	if !isDefault && !hasDefault && len(c.llms) == 0 {
		return fmt.Errorf("must have at least one default LLM")
	}

	// If len is greater than 0 and we have no default, we make the first one default
	if !isDefault && !hasDefault && len(c.llms) > 0 {
		for _, llm := range c.llms {
			llm.IsDefault = true
			break
		}
	}

	c.llms[provider] = &llmSettings{
		Provider:   provider,
		Credential: credential,
		Model:      model,
		IsDefault:  isDefault,
	}

	return nil
}

// LLM returns the configuration for a specific provider.
// Returns the LLM config and a boolean indicating if it was found.
func (c *Config) LLM(provider string) (*llmSettings, bool) {
	llm, ok := c.llms[provider]
	return llm, ok
}

// DefaultLLM returns the configuration for the default LLM provider.
// Returns the default LLM config and a boolean indicating if a default was found.
func (c *Config) DefaultLLM() (*llmSettings, bool) {
	for _, llm := range c.llms {
		if llm.IsDefault {
			return llm, true
		}
	}
	return nil, false
}

// CurrentModel returns the model of the default LLM provider.
func (c *Config) CurrentModel() (string, bool) {
	llm, ok := c.DefaultLLM()
	if !ok {
		return "", false
	}
	return llm.Model, true
}

func (c *Config) Marshal() ([]byte, error) {
	var configData struct {
		LLMS []*llmSettings `json:"llms"`
	}

	for _, llm := range c.llms {
		configData.LLMS = append(
			configData.LLMS, &llmSettings{
				Provider:   llm.Provider,
				Credential: llm.Credential,
				Model:      llm.Model,
				IsDefault:  llm.IsDefault,
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
	var configData struct {
		LLMS []*llmSettings `json:"llms"`
	}

	if err := json.Unmarshal(data, &configData); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	for _, data := range configData.LLMS {
		c.llms[data.Provider] = &llmSettings{
			Provider:   data.Provider,
			Credential: data.Credential,
			Model:      data.Model,
			IsDefault:  data.IsDefault,
		}
	}

	return nil
}

// Save writes the configuration to disk, validating the configuration before saving.
// Creates the configuration directory if it doesn't exist.
func (c *Config) Save() error {
	var errs []error

	hasUniqueDefault := false

	for _, llm := range c.llms {
		if llm.Provider == "" {
			errs = append(errs, errors.New("provider cannot be empty"))
		}

		if llm.Model == "" {
			errs = append(errs, errors.New("model cannot be empty"))
		}

		if llm.Credential == "" {
			errs = append(errs, errors.New("credential cannot be empty"))
		}

		if llm.IsDefault {
			if hasUniqueDefault {
				errs = append(errs, fmt.Errorf("multiple default LLMs found: %s", llm.Provider))
			} else {
				hasUniqueDefault = true
			}
		}
	}

	if !hasUniqueDefault {
		errs = append(errs, errors.New("no default LLM found"))
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

// Remove clears all LLM configurations and removes the config file from disk.
func (c *Config) Remove() error {
	c.llms = make(map[string]*llmSettings)

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

// Dir returns the path to the configuration directory.
func Dir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return path.Join(home, _dirName)
}
