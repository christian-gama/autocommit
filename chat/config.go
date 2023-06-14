package chat

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Config is the configuration for the chat service.
type Config struct {
	APIKey      string
	ModelName   string
	Template    string
	Temperature float32
}

// NewConfig creates a new chat service configuration.
func NewConfig(apiKey string, model string, verbose bool, temperature float32) *Config {
	kind, ok := ModelMap[model]
	if !ok {
		log.Fatalf("Model %s was not found", model)
	}

	template := fmt.Sprintf("%s\n%s", ShortMode, kind)
	if verbose {
		template = fmt.Sprintf("%s\n%s", DetailedMode, kind)
	}

	return &Config{
		APIKey:      apiKey,
		ModelName:   model,
		Template:    template,
		Temperature: temperature,
	}
}

func ValidateTemperature(temperature float32) error {
	if temperature <= 0 || temperature > 1 {
		return fmt.Errorf("temperature must be greater than 0 and less than or equal to 1")
	}
	return nil
}

func ValidateAPIKey(apiKey string) error {
	if apiKey == "" {
		return errors.New("API key cannot be empty")
	}

	if apiKey == "" {
		return errors.New("API key cannot be empty")
	}

	httpRequest, err := createRequestWithHeaders(http.MethodGet, openAIModelsURL, apiKey)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer httpResponse.Body.Close()

	if err := handleHTTPResponse(httpResponse); err != nil {
		return err
	}

	return nil
}

func ValidateModel(model string) error {
	if model == "" {
		return errors.New("model cannot be empty")
	}

	_, ok := ModelMap[model]
	if !ok {
		return fmt.Errorf("model %s was not found", model)
	}

	return nil
}
