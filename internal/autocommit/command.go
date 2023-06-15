package autocommit

import (
	"fmt"

	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/openai"
)

type GeneratorCommand interface {
	Execute() (string, error)
}

type generatorCommandImpl struct {
	chatCommand openai.ChatCommand
	diffCommand git.DiffCommand
}

func (g *generatorCommandImpl) Execute() (string, error) {
	diff, err := g.diffCommand.Execute()
	if err != nil {
		return "", err
	}

	system := openai.NewSystem(SystemMsg, "CommitMessageGenerator")

	response, err := g.chatCommand.Execute(system, fmt.Sprintf("Output of 'git diff':\n\"%s\"", diff))
	if err != nil {
		return "", err
	}

	return response, nil
}

func NewGeneratorCommand(chatCommand openai.ChatCommand, diffCommand git.DiffCommand) GeneratorCommand {
	return &generatorCommandImpl{
		chatCommand: chatCommand,
		diffCommand: diffCommand,
	}
}

type ClipboardCommand interface {
	Execute(message string) error
}

type clipboardCommandImpl struct {
	clipboard Clipboard
}

func (c *clipboardCommandImpl) Execute(message string) error {
	return c.clipboard.Copy(message)
}

func NewClipboardCommand(clipboard Clipboard) ClipboardCommand {
	return &clipboardCommandImpl{
		clipboard: clipboard,
	}
}
