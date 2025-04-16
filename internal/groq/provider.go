package groq

import (
	"fmt"

	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/storage"
)

type GroqProvider struct{}

func NewGroqProvider() *GroqProvider {
	return &GroqProvider{}
}

func (o *GroqProvider) MakeConfigRepo() llm.ConfigRepo {
	return llm.NewConfigRepo(storage.NewStorage("config.json"))
}

func (o *GroqProvider) ChatCommand() llm.ChatCommand {
	return llm.NewChatCommand(NewChat(o.MakeConfigRepo()))
}

func (o *GroqProvider) AskConfigsCli() llm.AskConfigsCli {
	return llm.NewAskConfigsCli(o)
}

func (o *GroqProvider) AskToChangeModelCli() llm.AskToChangeModelCli {
	return llm.NewAskToChangeModelCli()
}

func (o *GroqProvider) VerifyConfigCommand() llm.VerifyConfigCommand {
	return llm.NewVerifyConfigCommand(o.MakeConfigRepo())
}

func (o *GroqProvider) UpdateConfigCommand() llm.UpdateConfigCommand {
	return llm.NewUpdateConfigCommand(o.MakeConfigRepo())
}

func (o *GroqProvider) ResetConfigCommand() llm.ResetConfigCommand {
	return llm.NewResetConfigCommand(o.MakeConfigRepo())
}

func (o *GroqProvider) GetName() string {
	return "groq"
}

func (o *GroqProvider) GetValidationURL() string {
	return "https://api.groq.com/openai/v1/models"
}

func (o *GroqProvider) ValidateModel(model string) error {
	return ValidateModel(model)
}

func (o *GroqProvider) GetModelHelpText() string {
	return "The model to use for the Groq API."
}

func (o *GroqProvider) GetAPIKeyLabel() string {
	return "Groq API Key"
}

func (o *GroqProvider) GetAPIKeyHelpText() string {
	return "The Groq API Key is used to authenticate your requests to the Groq API."
}

func (o *GroqProvider) GetModelLabel() string {
	return "greq model"
}

func (o *GroqProvider) GetDefaultModel() string {
	return AllowedModels[1]
}

func (o *GroqProvider) ValidateTemperature(temperature float32) error {
	if temperature <= 0 || temperature > 1 {
		return fmt.Errorf("temperature must be greater than 0 and less than or equal to 1")
	}
	return nil
}

func (o *GroqProvider) GetAllowedModels() []string {
	return AllowedModels
}

func (o *GroqProvider) GetApiKeyHelpText() string {
	return "The Groq API Key is used to authenticate your requests to the Groq API."
}

func (o *GroqProvider) ValidateApiKey(apiKey string) error {
	return ValidateApiKey(apiKey)
}

func (o *GroqProvider) GetApiKeyLabel() string {
	return "Groq API Key"
}

func (o *GroqProvider) NewConfig(apiKey string, model string) llm.Config {
	return NewConfig(apiKey, model)
}
