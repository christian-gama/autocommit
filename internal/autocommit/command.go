package autocommit

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

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
	chatCommand   openai.ChatCommand
	diffCommand   git.DiffCommand
	systemMsgRepo SystemMsgRepo
}

// Execute implements the GeneratorCommand interface.
func (g *generatorCommandImpl) Execute(config *openai.Config) (string, error) {
	diff, err := g.diffCommand.Execute()
	if err != nil {
		return "", err
	}

	systemMsg, err := g.systemMsgRepo.GetSystemMsg()
	if err != nil {
		return "", err
	}

	system := openai.NewSystem(systemMsg, "CommitMessageGenerator")
	msg := fmt.Sprintf(
		"As a reminder, be concise and always write the texts in imperative mood and in present tense. Here is my 'git diff' output: \n\n%s",
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
	systemMsgRepo SystemMsgRepo,
) GeneratorCommand {
	return &generatorCommandImpl{
		chatCommand:   chatCommand,
		diffCommand:   diffCommand,
		systemMsgRepo: systemMsgRepo,
	}
}

// ClipboardCommand is the interface that wraps the Execute method.
type ClipboardCommand interface {
	Execute(message string) error
}

// clipboardCommandImpl is an implementation of ClipboardCommand.
type clipboardCommandImpl struct {
	clipboard Clipboard
}

// Execute implements the ClipboardCommand interface.
func (c *clipboardCommandImpl) Execute(message string) error {
	return c.clipboard.Copy(fmt.Sprintf("git commit -m \"%s\"", message))
}

// NewClipboardCommand creates a new instance of ClipboardCommand.
func NewClipboardCommand(clipboard Clipboard) ClipboardCommand {
	return &clipboardCommandImpl{
		clipboard: clipboard,
	}
}

// AddInstructionCommand is the interface that wraps the Execute method.
type AddInstructionCommand interface {
	Execute(config *openai.Config, instruction string) (string, error)
}

// addInstructionCommandImpl is an implementation of AddInstructionCommand.
type addInstructionCommandImpl struct {
	chatCommand openai.ChatCommand
}

// Execute implements the AddInstructionCommand interface.
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

// NewAddInstructionCommand creates a new instance of AddInstructionCommand.
func NewAddInstructionCommand(
	chatCommand openai.ChatCommand,
) AddInstructionCommand {
	return &addInstructionCommandImpl{
		chatCommand: chatCommand,
	}
}

// OpenSystemMsgCommand is the interface that wraps the Execute method.
type OpenSystemMsgCommand interface {
	Execute() error
}

// openSystemMsgCommandImpl is an implementation of OpenSystemMsgCommand.
type openSystemMsgCommandImpl struct{}

// Execute implements the OpenSystemMsgCommand interface.
func (u *openSystemMsgCommandImpl) Execute() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := path.Join(homeDir, ".autocommit", "system_msg.txt")

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", filePath)
	case "linux", "darwin":
		cmd = exec.Command("open", filePath)
	default:
		return errors.New("unsupported platform")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// NewOpenSystemMsgCommand creates a new instance of UpdateSystemMsgCommand.
func NewOpenSystemMsgCommand() OpenSystemMsgCommand {
	return &openSystemMsgCommandImpl{}
}

// SystemMsgHealthCheckCommand is the interface that wraps the Execute method.
type SystemMsgHealthCheckCommand interface {
	Execute() error
}

// systemMsgHealthCheckCommandImpl is an implementation of SystemMsgHealthCheckCommand.
type systemMsgHealthCheckCommandImpl struct {
	systemMsgRepo SystemMsgRepo
}

// Execute implements the SystemMsgHealthCheckCommand interface.
func (s *systemMsgHealthCheckCommandImpl) Execute() error {
	if !s.systemMsgRepo.Exists() {
		return s.systemMsgRepo.SaveSystemMsg()
	}

	if systemMsg, err := s.systemMsgRepo.GetSystemMsg(); err != nil {
		return err
	} else if systemMsg == "" {
		return s.systemMsgRepo.SaveSystemMsg()
	}

	return nil
}

// NewSystemMsgHealthCheckCommand creates a new instance of SystemMsgHealthCheckCommand.
func NewSystemMsgHealthCheckCommand(
	systemMsgRepo SystemMsgRepo,
) SystemMsgHealthCheckCommand {
	return &systemMsgHealthCheckCommandImpl{
		systemMsgRepo: systemMsgRepo,
	}
}
