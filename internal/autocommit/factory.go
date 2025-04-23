package autocommit

import (
	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/storage"
)

func MakeGeneratorCommand(provider llm.Provider) GeneratorCommand {
	chatCommand := provider.ChatCommand()
	diffCommand := git.MakeDiffCommand()

	return NewGeneratorCommand(chatCommand, diffCommand, MakeSystemMsgRepo())
}

func MakeClipboard() Clipboard {
	return NewClipboard()
}

func MakePostCommitCli() PostCommitCli {
	return NewPostCommitCli()
}

func MakeClipboardCommand() ClipboardCommand {
	return NewClipboardCommand(MakeClipboard())
}

func MakeAddInstructionCommand(provider llm.Provider) AddInstructionCommand {
	return NewAddInstructionCommand(provider.ChatCommand())
}

func MakeAddInstructionCli() AddInstructionCli {
	return NewAddInstructionCli()
}

func MakeSystemMsgRepo() SystemMsgRepo {
	return NewSystemMsgRepo(storage.NewStorage("system_msg.txt"))
}

func MakeSystemMsgHealthCheckCommand() SystemMsgHealthCheckCommand {
	return NewSystemMsgHealthCheckCommand(MakeSystemMsgRepo())
}

func MakeOpenSystemMsgCommand() OpenSystemMsgCommand {
	return NewOpenSystemMsgCommand()
}
