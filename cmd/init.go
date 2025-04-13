package cmd

import (
	"github.com/christian-gama/autocommit/internal/autocommit"
	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/openai"
	"github.com/christian-gama/autocommit/internal/provider"
)

var (
	postCommitCli         autocommit.PostCommitCli
	verifyConfigCommand   llm.VerifyConfigCommand
	generatorCommand      autocommit.GeneratorCommand
	askConfigsCli         llm.AskConfigsCli
	commitCommand         git.CommitCommand
	clipboardCommand      autocommit.ClipboardCommand
	resetConfigCommand    openai.ResetConfigCommand
	updateConfigCommand   llm.UpdateConfigCommand
	askToChangeModelCli   openai.AskToChangeModelCli
	addInstructionCommand autocommit.AddInstructionCommand
	addInstructionCli     autocommit.AddInstructionCli
	config                llm.Config
	openSystemMsgCommand  autocommit.OpenSystemMsgCommand
	systemMsgHealthCheck  autocommit.SystemMsgHealthCheckCommand
	llmProvider           llm.Provider
)

func init() {
	llmProvider = openai.NewOpenAIProvider()
	providerFactory := provider.NewProviderFactory(llmProvider)
	postCommitCli = autocommit.MakePostCommitCli()
	verifyConfigCommand = providerFactory.MakeVerifyConfigCommand()
	generatorCommand = autocommit.MakeGeneratorCommand(llmProvider)
	askConfigsCli = provider.MakeAskConfigsCli()
	commitCommand = git.MakeCommitCommand()
	clipboardCommand = autocommit.MakeClipboardCommand()
	resetConfigCommand = providerFactory.MakeResetConfigCommand()
	updateConfigCommand = providerFactory.MakeUpdateConfigCommand()
	askToChangeModelCli = providerFactory.MakeAskToChangeModelCli()
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
