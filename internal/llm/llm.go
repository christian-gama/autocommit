package llm

// Config is the configuration for any LLM provider.
type Config interface {
	// Provider returns the name of the LLM provider
	Provider() string
}

// System represents the system configuration for an LLM - typically the initial context/message
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

// Chat is the interface that wraps the Response method.
type Chat interface {
	// Response returns the response from the LLM.
	Response(config Config, system *System, input string) (string, error)
}

// ConfigRepo is the interface that wraps the basic operations with the config file.
type ConfigRepo interface {
	// SaveConfig saves the config.
	SaveConfig(config Config) error

	// GetConfig returns the config.
	GetConfig() (Config, error)

	// DeleteConfig deletes the config.
	DeleteConfig() error

	// UpdateConfig updates the config.
	UpdateConfig(config Config) error

	// Exists returns true if the config exists.
	Exists() bool
}

// AskConfigsCli is a command line interface that asks the user for the configuration.
type AskConfigsCli interface {
	Execute() (Config, error)
}

// AskToChangeModelCli asks the user if they want to change the model.
type AskToChangeModelCli interface {
	Execute() (bool, error)
}

// ChatCommand is the interface that wraps the basic Execute method.
type ChatCommand interface {
	// Execute returns the response from the LLM.
	Execute(config Config, system *System, input string) (string, error)
}

// VerifyConfigCommand is the interface that wraps the basic Execute method.
type VerifyConfigCommand interface {
	// Execute will verify if the configs were initialized and if not, it will initialize them.
	Execute(getConfigsFn func() (Config, error)) (Config, error)
}

// ResetConfigCommand is the interface that wraps the basic Execute method.
type ResetConfigCommand interface {
	// Execute will reset the configs.
	Execute() error
}

// UpdateConfigCommand is the interface that wraps the basic Execute method.
type UpdateConfigCommand interface {
	// Execute will update the configs.
	Execute(config Config) error
}
