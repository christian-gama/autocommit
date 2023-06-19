package autocommit

import (
	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/openai"
)

var chatCommand = openai.MakeChatCommand()

func MakeGeneratorCommand() GeneratorCommand {
	diffCommand := git.MakeDiffCommand()

	return NewGeneratorCommand(chatCommand, diffCommand)
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
