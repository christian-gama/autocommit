package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/christian-gama/autocommit/internal/autocommit"
	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/groq"
	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/openai"
)

var (
	postCommitCli         autocommit.PostCommitCli
	generatorCommand      autocommit.GeneratorCommand
	gitCommand            *git.Git
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
		return groq.NewGroqProvider()
	default:
		panic("Invalid choice")
	}
}

func init() {
	postCommitCli = autocommit.MakePostCommitCli()

	gitCommand = git.New()
	clipboardCommand = autocommit.MakeClipboardCommand()
	addInstructionCli = autocommit.MakeAddInstructionCli()
	openSystemMsgCommand = autocommit.MakeOpenSystemMsgCommand()
	systemMsgHealthCheck = autocommit.MakeSystemMsgHealthCheckCommand()
}

func init() {
	cmd.AddCommand(resetCmd)
	cmd.AddCommand(setCmd)
	cmd.AddCommand(editCmd)
}
