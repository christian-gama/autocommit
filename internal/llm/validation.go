package llm

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
)

func ValidateApiKey(apiKey string, provider Provider) error {
	if apiKey == "" {
		return errors.New("API key cannot be empty")
	}

	url := provider.GetValidationURL()

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

func ValidateModel(model string, provider Provider) error {
	if model == "" {
		return errors.New("model cannot be empty")
	}

	if slices.Contains(provider.GetAllowedModels(), model) {
		return nil
	}

	return fmt.Errorf("model %s was not found", model)
}
