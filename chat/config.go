package chat

import "fmt"

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
		panic(fmt.Sprintf("Model %s was not found", model))
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
