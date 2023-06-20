package autocommit

import (
	"fmt"

	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/openai"
)

// GeneratorCommand is the interface that wraps the Execute method.
type GeneratorCommand interface {
	// Execute returns the commit message generated by OpenAI.
	Execute(config *openai.Config) (string, error)
}

// generatorCommandImpl is an implementation of GeneratorCommand.
type generatorCommandImpl struct {
	chatCommand openai.ChatCommand
	diffCommand git.DiffCommand
}

// Execute implements the GeneratorCommand interface.
func (g *generatorCommandImpl) Execute(config *openai.Config) (string, error) {

	diff, err := g.diffCommand.Execute()
	if err != nil {
		return "", err
	}

	system := openai.NewSystem(SystemMsg, "CommitMessageGenerator")
	msg := fmt.Sprintf(
		"Create a commit message based on the output of 'git diff' below. As a reminder, be concise and always write the texts in imperative mood and in present tense:\n\n%s",
		diff,
	)

	fmt.Printf("⌛ Creating a commit message...\n")
	response, err := g.chatCommand.Execute(
		config,
		system,
		msg,
	)
	if err != nil {
		return "", err
	}

	return response, nil
}

// NewGeneratorCommand creates a new instance of GeneratorCommand.
func NewGeneratorCommand(
	chatCommand openai.ChatCommand,
	diffCommand git.DiffCommand,
) GeneratorCommand {
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
	return c.clipboard.Copy(fmt.Sprintf("git commit -m \"%s\"", message))
}

func NewClipboardCommand(clipboard Clipboard) ClipboardCommand {
	return &clipboardCommandImpl{
		clipboard: clipboard,
	}
}

type AddInstructionCommand interface {
	Execute(config *openai.Config, instruction string) (string, error)
}

type addInstructionCommandImpl struct {
	chatCommand openai.ChatCommand
}

func (a *addInstructionCommandImpl) Execute(
	config *openai.Config,
	instruction string,
) (string, error) {
	fmt.Printf("💡 Enhancing the message with your new instruction...\n")
	response, err := a.chatCommand.Execute(config, nil, instruction)
	if err != nil {
		return "", err
	}

	return response, nil
}

func NewAddInstructionCommand(
	chatCommand openai.ChatCommand,
) AddInstructionCommand {
	return &addInstructionCommandImpl{
		chatCommand: chatCommand,
	}
}
