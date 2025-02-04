package openai

// Config is the configuration for the OpenAI API.
type Config struct {
	ApiKey      string  `json:"openai_api_key"`
	Temperature float32 `json:"openai_temperature"`
	Model       string  `json:"openai_model"`
}

// NewConfig creates a new Config.
func NewConfig(apiKey, model string) *Config {
	return &Config{
		ApiKey:      apiKey,
		Temperature: 0.3,
		Model:       model,
	}
}

// System is the system configuration for the OpenAI API - the initial message and name.
type System struct {
	Message string
	Name    string
}

// NewSystem creates a new System.
func NewSystem(message, name string) *System {
	return &System{
		Message: message,
		Name:    name,
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
