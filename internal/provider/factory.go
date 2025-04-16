package provider

import (
	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/storage"
)

type ProviderFactory struct {
	provider llm.Provider
}

func NewProviderFactory(provider llm.Provider) *ProviderFactory {
	return &ProviderFactory{provider: provider}
}

func (f *ProviderFactory) MakeConfigRepo() llm.ConfigRepo {
	return llm.NewConfigRepo(storage.NewStorage("config.json"))
}

func (f *ProviderFactory) MakeVerifyConfigCommand() llm.VerifyConfigCommand {
	return llm.NewVerifyConfigCommand(f.MakeConfigRepo())
}

func (f *ProviderFactory) MakeAskToChangeModelCli() llm.AskToChangeModelCli {
	return llm.NewAskToChangeModelCli()
}

func (f *ProviderFactory) MakeUpdateConfigCommand() llm.UpdateConfigCommand {
	return llm.NewUpdateConfigCommand(f.MakeConfigRepo())
}

func (f *ProviderFactory) MakeResetConfigCommand() llm.ResetConfigCommand {
	return llm.NewResetConfigCommand(f.MakeConfigRepo())
}

func (f *ProviderFactory) MakeAskConfigsCli() llm.AskConfigsCli {
	return llm.NewAskConfigsCli(f.provider)
}
