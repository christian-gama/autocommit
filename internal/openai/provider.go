package openai

import (
	"fmt"

	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/storage"
)

type OpenAIProvider struct{}

func NewOpenAIProvider() *OpenAIProvider {
	return &OpenAIProvider{}
}

func (o *OpenAIProvider) MakeConfigRepo() llm.ConfigRepo {
	return NewConfigRepo(storage.NewStorage("config.json"))
}

func (o *OpenAIProvider) ChatCommand() llm.ChatCommand {
	return NewChatCommand(NewChat(o.MakeConfigRepo()))
}

func (o *OpenAIProvider) AskConfigsCli() llm.AskConfigsCli {
	return llm.NewAskConfigsCli(o)
}

func (o *OpenAIProvider) AskToChangeModelCli() llm.AskToChangeModelCli {
	return llm.NewAskToChangeModelCli()
}

func (o *OpenAIProvider) VerifyConfigCommand() llm.VerifyConfigCommand {
	return llm.NewVerifyConfigCommand(o.MakeConfigRepo())
}

func (o *OpenAIProvider) UpdateConfigCommand() llm.UpdateConfigCommand {
	return llm.NewUpdateConfigCommand(o.MakeConfigRepo())
}

func (o *OpenAIProvider) ResetConfigCommand() llm.ResetConfigCommand {
	return llm.NewResetConfigCommand(o.MakeConfigRepo())
}

func (o *OpenAIProvider) GetName() string {
	return "openai"
}

func (o *OpenAIProvider) GetValidationURL() string {
	return "https://platform.openai.com/docs/models/gpt"
}

func (o *OpenAIProvider) ValidateModel(model string) error {
	return ValidateModel(model)
}

func (o *OpenAIProvider) GetModelHelpText() string {
	return "The model to use for the OpenAI API."
}

func (o *OpenAIProvider) GetAPIKeyLabel() string {
	return "OpenAI API Key"
}

func (o *OpenAIProvider) GetAPIKeyHelpText() string {
	return "The OpenAI API Key is used to authenticate your requests to the OpenAI API."
}

func (o *OpenAIProvider) GetModelLabel() string {
	return "openai model"
}

func (o *OpenAIProvider) GetDefaultModel() string {
	return AllowedModels[0]
}

func (o *OpenAIProvider) ValidateTemperature(temperature float32) error {
	if temperature <= 0 || temperature > 1 {
		return fmt.Errorf("temperature must be greater than 0 and less than or equal to 1")
	}
	return nil
}

func (o *OpenAIProvider) GetAllowedModels() []string {
	return AllowedModels
}

func (o *OpenAIProvider) GetApiKeyHelpText() string {
	return "The OpenAI API Key is used to authenticate your requests to the OpenAI API."
}

func (o *OpenAIProvider) ValidateApiKey(apiKey string) error {
	return ValidateApiKey(apiKey)
}

func (o *OpenAIProvider) GetApiKeyLabel() string {
	return "OpenAI API Key"
}

func (o *OpenAIProvider) NewConfig(apiKey string, model string) llm.Config {
	return NewConfig(apiKey, model)
}
