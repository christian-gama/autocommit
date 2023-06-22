package autocommit

import (
	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/openai"
	"github.com/christian-gama/autocommit/internal/storage"
)

var chatCommand = openai.MakeChatCommand()

func MakeGeneratorCommand() GeneratorCommand {
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

func MakeAddInstructionCommand() AddInstructionCommand {
	return NewAddInstructionCommand(chatCommand)
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
