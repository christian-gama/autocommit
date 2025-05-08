package groq

type GroqConfig struct {
	ApiKey      string  `json:"groq_api_key"`
	Temperature float32 `json:"groq_temperature"`
	Model       string  `json:"groq_model"`
	
	apiKeySet      bool `json:"-"`
	temperatureSet bool `json:"-"`
	modelSet       bool `json:"-"`
}

// NewGroqConfig creates a new Config.
func NewGroqConfig(apiKey, model string, temperature float32) *GroqConfig {
	return &GroqConfig{
		ApiKey:      apiKey,
		Temperature: temperature,
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
	c.apiKeySet = true
}

func (c *GroqConfig) SetTemperature(temperature float32) {
	c.Temperature = temperature
	c.temperatureSet = true
}

func (c *GroqConfig) SetModel(model string) {
	c.Model = model
	c.modelSet = true
}

func (c *GroqConfig) IsAPIKeySet() bool {
	return c.apiKeySet
}

func (c *GroqConfig) IsTemperatureSet() bool {
	return c.temperatureSet
}

func (c *GroqConfig) IsModelSet() bool {
	return c.modelSet
}

func (c *GroqConfig) MarkAPIKeySet() {
	c.apiKeySet = true
}

func (c *GroqConfig) MarkModelSet() {
	c.modelSet = true
}

func (c *GroqConfig) MarkTemperatureSet() {
	c.temperatureSet = true
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
