package openai

import (
	"errors"

	"github.com/christian-gama/autocommit/internal/storage"
)

type Repo interface {
	SaveConfig(config *Config) error
	GetConfig() (*Config, error)
	DeleteConfig() error
	UpdateConfig(config *Config) error
	Exists() bool
}

type repoImpl struct {
	storage *storage.Storage
}

func (r *repoImpl) DeleteConfig() error {
	return r.storage.Delete()
}

func (r *repoImpl) GetConfig() (*Config, error) {
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

func (r *repoImpl) SaveConfig(config *Config) error {
	content, err := MarshalConfig(config)
	if err != nil {
		return err
	}

	return r.storage.Create(content)
}

func (r *repoImpl) UpdateConfig(config *Config) error {
	content, err := MarshalConfig(config)
	if err != nil {
		return err
	}

	return r.storage.Update(content)
}

func (r *repoImpl) Exists() bool {
	content, err := r.storage.Read()
	if err != nil {
		return false
	}

	return len(content) > 0
}

func NewRepo(storage *storage.Storage) Repo {
	return &repoImpl{storage: storage}
}
