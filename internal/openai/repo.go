package openai

import (
	"errors"

	"github.com/christian-gama/autocommit/internal/storage"
)

// ConfigRepo is the interface that wraps the basic operations with the config file.
type ConfigRepo interface {
	// SaveConfig saves the config file.
	SaveConfig(config *Config) error

	// GetConfig returns the config file.
	GetConfig() (*Config, error)

	// DeleteConfig deletes the config file.
	DeleteConfig() error

	// UpdateConfig updates the config file.
	UpdateConfig(config *Config) error

	// Exists returns true if the config file exists.
	Exists() bool
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
func (r *configRepoImpl) GetConfig() (*Config, error) {
	content, err := r.storage.Read()
	if err != nil {
		return nil, err
	}

	config, err := UnmarshalConfig(content)
	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, errors.New("Config is empty")
	}

	return config, nil
}

// SaveConfig implements the Repo interface.
func (r *configRepoImpl) SaveConfig(config *Config) error {
	content, err := MarshalConfig(config)
	if err != nil {
		return err
	}

	return r.storage.Create(content)
}

// UpdateConfig implements the Repo interface.
func (r *configRepoImpl) UpdateConfig(config *Config) error {
	content, err := MarshalConfig(config)
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

// NewConfigRepo creates a new instance of Repo.
func NewConfigRepo(storage *storage.Storage) ConfigRepo {
	return &configRepoImpl{storage: storage}
}
