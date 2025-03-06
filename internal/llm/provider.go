package llm

type Provider interface {
	GetName() string
	GetAllowedModels() []string
	AskConfigsCli() AskConfigsCli
	AskToChangeModelCli() AskToChangeModelCli
	VerifyConfigCommand() VerifyConfigCommand
	ResetConfigCommand() ResetConfigCommand
	UpdateConfigCommand() UpdateConfigCommand
	ChatCommand() ChatCommand
	GetValidationURL() string
	ValidateModel(string) error
	ValidateTemperature(float32) error
}
