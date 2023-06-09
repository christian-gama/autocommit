package chat

// Config is the configuration for the chat service.
type Config struct {
	APIKey      string
	Model       string
	Template    string
	Temperature float32
}

// NewConfig creates a new chat service configuration.
func NewConfig(apiKey string, model string, verbose bool, temperature float32) *Config {
	template := ShortTemplate
	if verbose {
		template = VerboseTemplate
	}

	return &Config{
		APIKey:      apiKey,
		Model:       model,
		Template:    template,
		Temperature: temperature,
	}
}
