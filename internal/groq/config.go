package groq

type GroqConfig struct {
	ApiKey      string  `json:"groq_api_key"`
	Temperature float32 `json:"groq_temperature"`
	Model       string  `json:"groq_model"`
}

// NewConfig creates a new Config.
func NewConfig(apiKey, model string) *GroqConfig {
	return &GroqConfig{
		ApiKey:      apiKey,
		Temperature: 0.3,
		Model:       model,
	}
}

func (c *GroqConfig) GetName() string {
	return "groq"
}

func (c *GroqConfig) GetAPIKey() string {
	return c.ApiKey
}

func (c *GroqConfig) GetTemperature() float32 {
	return c.Temperature
}

func (c *GroqConfig) GetModel() string {
	return c.Model
}

func (c *GroqConfig) SetAPIKey(ApiKey string) {
	c.ApiKey = ApiKey
}

func (c *GroqConfig) SetTemperature(temperature float32) {
	c.Temperature = temperature
}

func (c *GroqConfig) SetModel(model string) {
	c.Model = model
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
	DEEPSEEK = "deepseek-r1-distill-llama-70b"
	LLAMA    = "llama3-70b-8192"
)

var AllowedModels = []string{
	DEEPSEEK,
	LLAMA,
}
