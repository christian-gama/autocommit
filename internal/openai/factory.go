package openai

import "github.com/christian-gama/autocommit/internal/storage"

func MakeConfigRepo() Repo {
	return NewRepo(storage.NewStorage("config.json"))
}

func MakeAskConfigsCli() AskConfigsCli {
	return NewAskConfigsCli()
}

func MakeVerifyConfigCommand() VerifyConfigCommand {
	return NewVerifyConfigCommand(MakeConfigRepo())
}

func MakeResetConfigCommand() ResetConfigCommand {
	return NewResetConfigCommand(MakeConfigRepo())
}

func MakeUpdateConfigCommand() UpdateConfigCommand {
	return NewUpdateConfigCommand(MakeConfigRepo())
}

func MakeAskToChangeModelCli() AskToChangeModelCli {
	return NewAskToChangeModelCli()
}
