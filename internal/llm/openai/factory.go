package openai

import (
	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/storage"
)

// factory implements the llm.Factory interface for OpenAI
type factory struct{}

func (f *factory) MakeConfigRepo() llm.ConfigRepo {
	return NewConfigRepo(storage.NewStorage("config.json"))
}

func (f *factory) MakeAskConfigsCli() llm.AskConfigsCli {
	return NewAskConfigsCli()
}

func (f *factory) MakeVerifyConfigCommand() llm.VerifyConfigCommand {
	return llm.NewVerifyConfigCommand(f.MakeConfigRepo())
}

func (f *factory) MakeResetConfigCommand() llm.ResetConfigCommand {
	return llm.NewResetConfigCommand(f.MakeConfigRepo())
}

func (f *factory) MakeUpdateConfigCommand() llm.UpdateConfigCommand {
	return newUpdateConfigCommand(f.MakeConfigRepo())
}

func (f *factory) MakeAskToChangeModelCli() llm.AskToChangeModelCli {
	return NewAskToChangeModelCli()
}

func (f *factory) MakeChatCommand() llm.ChatCommand {
	return llm.NewChatCommand(NewChat(f.MakeConfigRepo().(ConfigRepo)))
}

// NewFactory creates an OpenAI implementation of the LLM factory
func NewFactory() llm.Factory {
	return &factory{}
}
