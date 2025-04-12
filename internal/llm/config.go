package llm

type Config interface {
	GetName() string
	GetAPIKey() string
	GetTemperature() float32
	GetModel() string

	SetAPIKey(string)
	SetTemperature(float32)
	SetModel(string)
}

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
