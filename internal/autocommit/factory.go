package autocommit

import (
	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/openai"
)

func MakeGeneratorCommand() GeneratorCommand {
	chatCommand := openai.NewChatCommand(openai.NewChat(openai.MakeConfigRepo()))
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
