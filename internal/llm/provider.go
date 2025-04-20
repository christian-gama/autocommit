package llm

// Provider is the interface for LLM providers
type Provider interface {
	GetName() string
	GetConfigFileName() string
	GetAllowedModels() []string
	AskConfigsCli() AskConfigsCli
	AskToChangeModelCli() AskToChangeModelCli
	VerifyConfigCommand() VerifyConfigCommand
	ResetConfigCommand() ResetConfigCommand
	UpdateConfigCommand() UpdateConfigCommand
	ChatCommand() ChatCommand
	GetValidationURL() string
	GetDefaultModel() string
	GetModelHelpText() string
	GetApiKeyLabel() string
	GetApiKeyHelpText() string
	NewConfig(apiKey string, model string) Config
	ValidateModel(string) error
	ValidateTemperature(float32) error
	MarshalConfig(Config) ([]byte, error)
	UnmarshalConfig([]byte) (Config, error)
}
