package openai

// Config is the configuration for the OpenAI API.
type Config struct {
	ApiKey      string  `json:"openai_api_key"`
	Temperature float32 `json:"openai_temperature"`
	Model       string  `json:"openai_model"`
}

// Provider returns the name of the LLM provider
func (c *Config) Provider() string {
	return "openai"
}

// NewConfig creates a new Config.
func NewConfig(apiKey, model string) *Config {
	return &Config{
		ApiKey:      apiKey,
		Temperature: 0.3,
		Model:       model,
	}
}

const (
	O1Mini      = "o1-mini"
	GPT4o       = "gpt-4o"
	GPT4oLatest = "chatgpt-4o-latest"
	GPT4oMini   = "gpt-4o-mini"
	GPT4Turbo   = "gpt-4-turbo"
)

var AllowedModels = []string{
	O1Mini,
	GPT4o,
	GPT4oLatest,
	GPT4oMini,
	GPT4Turbo,
}
