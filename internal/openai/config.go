package openai

// *OpenAIConfig is the configuration for the OpenAI API.
type OpenAIConfig struct {
	ApiKey      string  `json:"openai_api_key"`
	Temperature float32 `json:"openai_temperature"`
	Model       string  `json:"openai_model"`
	
	apiKeySet      bool `json:"-"`
	temperatureSet bool `json:"-"`
	modelSet       bool `json:"-"`
}

// NewOpenAIConfig creates a new Config.
func NewOpenAIConfig(apiKey, model string, temperature float32) *OpenAIConfig {
	return &OpenAIConfig{
		ApiKey:      apiKey,
		Temperature: temperature,
		Model:       model,
	}
}

func (c *OpenAIConfig) GetName() string {
	return "openai"
}

func (c *OpenAIConfig) GetAPIKey() string {
	return c.ApiKey
}

func (c *OpenAIConfig) GetTemperature() float32 {
	return c.Temperature
}

func (c *OpenAIConfig) GetModel() string {
	return c.Model
}

func (c *OpenAIConfig) SetAPIKey(ApiKey string) {
	c.ApiKey = ApiKey
	c.apiKeySet = true
}

func (c *OpenAIConfig) SetTemperature(temperature float32) {
	c.Temperature = temperature
	c.temperatureSet = true
}

func (c *OpenAIConfig) SetModel(model string) {
	c.Model = model
	c.modelSet = true
}

func (c *OpenAIConfig) IsAPIKeySet() bool {
	return c.apiKeySet
}

func (c *OpenAIConfig) IsTemperatureSet() bool {
	return c.temperatureSet
}

func (c *OpenAIConfig) IsModelSet() bool {
	return c.modelSet
}

func (c *OpenAIConfig) MarkAPIKeySet() {
	c.apiKeySet = true
}

func (c *OpenAIConfig) MarkModelSet() {
	c.modelSet = true
}

func (c *OpenAIConfig) MarkTemperatureSet() {
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
