package cmd

import (
	"github.com/christian-gama/autocommit/internal/autocommit"
	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/llm"
	"github.com/christian-gama/autocommit/internal/llm/openai"
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
	factory               llm.Factory
)

func init() {
	factory = openai.NewFactory()

	postCommitCli = autocommit.MakePostCommitCli()
	verifyConfigCommand = factory.MakeVerifyConfigCommand()
	generatorCommand = autocommit.MakeGeneratorCommand()
	askConfigsCli = factory.MakeAskConfigsCli()
	commitCommand = git.MakeCommitCommand()
	clipboardCommand = autocommit.MakeClipboardCommand()
	resetConfigCommand = factory.MakeResetConfigCommand()
	updateConfigCommand = factory.MakeUpdateConfigCommand()
	askToChangeModelCli = factory.MakeAskToChangeModelCli()
	addInstructionCommand = autocommit.MakeAddInstructionCommand()
	addInstructionCli = autocommit.MakeAddInstructionCli()
	openSystemMsgCommand = autocommit.MakeOpenSystemMsgCommand()
	systemMsgHealthCheck = autocommit.MakeSystemMsgHealthCheckCommand()
}

func init() {
	cmd.AddCommand(resetCmd)
	cmd.AddCommand(setCmd)
	cmd.AddCommand(editCmd)
}
