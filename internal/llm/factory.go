package llm

// Factory creates implementations of the LLM interfaces
type Factory interface {
	MakeConfigRepo() ConfigRepo
	MakeAskConfigsCli() AskConfigsCli
	MakeVerifyConfigCommand() VerifyConfigCommand
	MakeResetConfigCommand() ResetConfigCommand
	MakeUpdateConfigCommand() UpdateConfigCommand
	MakeAskToChangeModelCli() AskToChangeModelCli
	MakeChatCommand() ChatCommand
}
