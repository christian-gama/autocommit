package llm

// Config is the interface that wraps the basic operations with the config file.
type Config interface {
	GetName() string
	GetAPIKey() string
	GetTemperature() float32
	GetModel() string
	SetAPIKey(string)
	SetTemperature(float32)
	SetModel(string)
	IsAPIKeySet() bool
	IsModelSet() bool
	IsTemperatureSet() bool
	MarkAPIKeySet()
	MarkModelSet()
	MarkTemperatureSet()
}

// ConfigImpl is an implementation of Config
type ConfigImpl struct {
	ApiKey      string  `json:"api_key"`
	Temperature float32 `json:"temperature"`
	Model       string  `json:"model"`
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
