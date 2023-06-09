package store

import (
	"encoding/json"
	"os"
	"path"
)

// Config represents the configuration of the application.
type Config struct {
	OpenAIAPIKey      string  `json:"openai_api_key"`
	OpenAIModel       string  `json:"openai_model"`
	OpenAITemperature float32 `json:"openai_temperature"`
}

// ConfigStore is the store for the configuration of the application.
type ConfigStore struct{}

// NewStore creates a new store for the configuration of the application.
func NewStore() *ConfigStore {
	return &ConfigStore{}
}

// IsStored returns true if the configuration file exists.
func (s *ConfigStore) IsStored() bool {
	configPath := getConfigPath()
	_, err := os.Stat(configPath)
	return !os.IsNotExist(err)
}

// Config returns the configuration of the application.
func (s *ConfigStore) Config() *Config {
	configFile, err := os.Open(getConfigPath())
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		panic(err)
	}

	return &config
}

// CreateConfigFile creates the configuration file for the application.
func (s *ConfigStore) CreateConfigFile(config *Config) {
	configPath := getAutocommitPath()

	// Create directory if it doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.Mkdir(configPath, 0755); err != nil {
			panic(err)
		}
	}

	configFile, err := os.Create(getConfigPath())
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	jsonParser := json.NewEncoder(configFile)
	if err := jsonParser.Encode(config); err != nil {
		panic(err)
	}
}

func getAutocommitPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(home, ".autocommit")
}

func getConfigPath() string {
	return path.Join(getAutocommitPath(), "config.json")
}
