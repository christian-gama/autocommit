package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

const (
	ConfigFileName    = "config.json"
	ConfigFileDirName = ".autocommit"
)

// Config represents the configuration of the application.
type Config struct {
	OpenAIAPIKey      string  `json:"openai_api_key"`
	OpenAIModel       string  `json:"openai_model"`
	OpenAITemperature float32 `json:"openai_temperature"`
}

// Store is the store for the configuration of the application.
type Store struct{}

// NewStore creates a new store for the configuration of the application.
func NewStore() *Store {
	return &Store{}
}

// IsStored returns true if the configuration file exists.
func (s *Store) IsStored() bool {
	configPath := getConfigPath()
	_, err := os.Stat(configPath)
	return !os.IsNotExist(err)
}

// Config returns the configuration of the application.
func (s *Store) Config() *Config {
	configFile, err := os.Open(getConfigPath())
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}

// CreateConfigFile creates the configuration file for the application.
func (s *Store) CreateConfigFile(config *Config) {
	configPath := getAutocommitPath()

	// Create directory if it doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.Mkdir(configPath, 0755); err != nil {
			log.Fatal(err)
		}
	}

	configFile, err := os.Create(getConfigPath())
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	jsonParser := json.NewEncoder(configFile)
	if err := jsonParser.Encode(config); err != nil {
		log.Fatal(err)
	}
}

// DeleteConfigFile deletes the configuration file for the application.
func (s *Store) DeleteConfigFile() {
	configPath := getConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Configuration file does not exist.")
		return
	}

	fmt.Println("Resetting configuration file...")
	if err := os.Remove(configPath); err != nil {
		log.Fatal(err)
	}
	fmt.Println(
		"Configuration file reset successfully - Next time you run autocommit, you will be asked to configure it again.",
	)
}

// SetOpenAIAPIKey sets the OpenAI API key in the configuration file.
func (s *Store) SetOpenAIAPIKey(apiKey string) {
	config := s.Config()
	config.OpenAIAPIKey = apiKey
	s.CreateConfigFile(config)
}

// SetOpenAIModel sets the OpenAI model in the configuration file.
func (s *Store) SetOpenAIModel(model string) {
	config := s.Config()
	config.OpenAIModel = model
	s.CreateConfigFile(config)
}

// SetOpenAITemperature sets the OpenAI temperature in the configuration file.
func (s *Store) SetOpenAITemperature(temperature float32) {
	config := s.Config()
	config.OpenAITemperature = temperature
	s.CreateConfigFile(config)
}

// LoadConfig loads the configuration of the application.
func Load() *Config {
	configStore := NewStore()

	if !configStore.IsStored() {
		configAnswers := askUserForConfig()
		configStore.CreateConfigFile(&Config{
			OpenAIAPIKey:      configAnswers.OpenAIAPIKey,
			OpenAIModel:       configAnswers.Model,
			OpenAITemperature: configAnswers.Temperature,
		})
	}

	return configStore.Config()
}

func getAutocommitPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(home, ConfigFileDirName)
}

func getConfigPath() string {
	return path.Join(getAutocommitPath(), ConfigFileName)
}
