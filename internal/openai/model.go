package openai

type Config struct {
	ApiKey      string  `json:"openai_api_key"`
	Temperature float32 `json:"openai_temperature"`
	Model       string  `json:"openai_model"`
}

func NewConfig(apiKey, model string, temperature float32) *Config {
	return &Config{
		ApiKey:      apiKey,
		Temperature: temperature,
		Model:       model,
	}
}

type System struct {
	Message string
	Name    string
}

func NewSystem(message, name string) *System {
	return &System{
		Message: message,
		Name:    name,
	}

}

const (
	GPT3Dot5Turbo16k = "gpt-3.5-turbo-16k"
	GPT3Dot5Turbo    = "gpt-3.5-turbo"
	GPT4             = "gpt-4"
	GPT432K          = "gpt-4-32k"
)

var Models = []string{
	GPT3Dot5Turbo,
	GPT3Dot5Turbo16k,
	GPT4,
	GPT432K,
}
