package openai

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
)

// ValidateTemperature validates the temperature for the OpenAI API.
func ValidateTemperature(temperature float32) error {
	if temperature <= 0 || temperature > 1 {
		return fmt.Errorf(
			"temperature must be greater than 0 and less than or equal to 1",
		)
	}
	return nil
}

// ValidateApiKey validates the API key for the OpenAI API. It does so
// by making a request to the models endpoint - if it fails, the API key is invalid.
func ValidateApiKey(apiKey string) error {
	if apiKey == "" {
		return errors.New("API key cannot be empty")
	}

	url := "https://api.openai.com/v1/models"
	if apiKey == "" {
		return errors.New("API key cannot be empty")
	}

	httpRequest, err := httpRequest(http.MethodGet, url, apiKey)
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

// ValidateModel validates the model for the OpenAI API by checking if the model is
// in the list of allowed models.
func ValidateModel(model string) error {
	if model == "" {
		return errors.New("model cannot be empty")
	}

	if slices.Contains(AllowedModels, model) {
		return nil
	}

	return fmt.Errorf("model %s was not found", model)
}
