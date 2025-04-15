package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/christian-gama/autocommit/internal/autocommit"
	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/openai"
)

var (
	postCommitCli         autocommit.PostCommitCli
	verifyConfigCommand   llm.VerifyConfigCommand
	generatorCommand      autocommit.GeneratorCommand
	askConfigsCli         llm.AskConfigsCli
	commitCommand         git.CommitCommand
	clipboardCommand      autocommit.ClipboardCommand
	resetConfigCommand    llm.ResetConfigCommand
	updateConfigCommand   llm.UpdateConfigCommand
	askToChangeModelCli   llm.AskToChangeModelCli
	addInstructionCommand autocommit.AddInstructionCommand
	addInstructionCli     autocommit.AddInstructionCli
	config                llm.Config
	openSystemMsgCommand  autocommit.OpenSystemMsgCommand
	systemMsgHealthCheck  autocommit.SystemMsgHealthCheckCommand
	llmProvider           llm.Provider
)

func loadProvider() llm.Provider {
	var choice string
	_ = survey.AskOne(&survey.Select{
		Message: "Choose your LLM Provider",
		Options: []string{"OpenAI", "Groq"},
	}, &choice)

	switch choice {
	case "OpenAI":
		return openai.NewOpenAIProvider()
	case "Groq":
		fmt.Println("Groq not implemented yet")
		os.Exit(0)
	default:
		panic("Invalid choice")
	}
	return nil
}

func init() {
	llmProvider = openai.NewOpenAIProvider()
	postCommitCli = autocommit.MakePostCommitCli()
	generatorCommand = autocommit.MakeGeneratorCommand(llmProvider)
	commitCommand = git.MakeCommitCommand()
	clipboardCommand = autocommit.MakeClipboardCommand()
	addInstructionCommand = autocommit.MakeAddInstructionCommand(llmProvider)
	addInstructionCli = autocommit.MakeAddInstructionCli()
	openSystemMsgCommand = autocommit.MakeOpenSystemMsgCommand()
	systemMsgHealthCheck = autocommit.MakeSystemMsgHealthCheckCommand()
}

func init() {
	cmd.AddCommand(resetCmd)
	cmd.AddCommand(setCmd)
	cmd.AddCommand(editCmd)
}
