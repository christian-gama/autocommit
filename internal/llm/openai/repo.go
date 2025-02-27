package openai

import (
	"encoding/json"
	"errors"

	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/storage"
)

// ConfigRepo is the interface that wraps the basic operations with the config file.
type ConfigRepo interface {
	llm.ConfigRepo
}

// configRepoImpl is an implementation of Repo.
type configRepoImpl struct {
	storage *storage.Storage
}

// DeleteConfig implements the Repo interface.
func (r *configRepoImpl) DeleteConfig() error {
	return r.storage.Delete()
}

// GetConfig implements the Repo interface.
func (r *configRepoImpl) GetConfig() (llm.Config, error) {
	content, err := r.storage.Read()
	if err != nil {
		return nil, err
	}

	config, err := unmarshalConfig(content)
	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, errors.New("Config is empty")
	}

	return config, nil
}

// SaveConfig implements the Repo interface.
func (r *configRepoImpl) SaveConfig(config llm.Config) error {
	openAIConfig, ok := config.(*Config)
	if !ok {
		return errors.New("invalid config type: expected OpenAI config")
	}

	content, err := marshalConfig(openAIConfig)
	if err != nil {
		return err
	}

	return r.storage.Create(content)
}

// UpdateConfig implements the Repo interface.
func (r *configRepoImpl) UpdateConfig(config llm.Config) error {
	openAIConfig, ok := config.(*Config)
	if !ok {
		return errors.New("invalid config type: expected OpenAI config")
	}

	content, err := marshalConfig(openAIConfig)
	if err != nil {
		return err
	}

	return r.storage.Update(content)
}

// Exists implements the Repo interface.
func (r *configRepoImpl) Exists() bool {
	content, err := r.storage.Read()
	if err != nil {
		return false
	}

	return len(content) > 0
}

// unmarshalConfig unmarshals a Config from a byte slice.
func unmarshalConfig(data []byte) (*Config, error) {
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// marshalConfig marshals a Config into a byte slice.
func marshalConfig(config *Config) ([]byte, error) {
	return json.Marshal(config)
}

// NewConfigRepo creates a new instance of Repo.
func NewConfigRepo(storage *storage.Storage) ConfigRepo {
	return &configRepoImpl{storage: storage}
}
