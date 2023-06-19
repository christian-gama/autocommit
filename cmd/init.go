package cmd

import (
	"github.com/christian-gama/autocommit/internal/autocommit"
	"github.com/christian-gama/autocommit/internal/git"
	"github.com/christian-gama/autocommit/internal/openai"
)

var (
	postCommitCli       autocommit.PostCommitCli
	verifyConfigCommand openai.VerifyConfigCommand
	generatorCommand    autocommit.GeneratorCommand
	askConfigsCli       openai.AskConfigsCli
	commitCommand       git.CommitCommand
	clipboardCommand    autocommit.ClipboardCommand
	resetConfigCommand  openai.ResetConfigCommand
	updateConfigCommand openai.UpdateConfigCommand
	askToChangeModelCli openai.AskToChangeModelCli
	config              *openai.Config
)

func init() {
	postCommitCli = autocommit.MakePostCommitCli()
	verifyConfigCommand = openai.MakeVerifyConfigCommand()
	generatorCommand = autocommit.MakeGeneratorCommand()
	askConfigsCli = openai.MakeAskConfigsCli()
	commitCommand = git.MakeCommitCommand()
	clipboardCommand = autocommit.MakeClipboardCommand()
	resetConfigCommand = openai.MakeResetConfigCommand()
	updateConfigCommand = openai.MakeUpdateConfigCommand()
	askToChangeModelCli = openai.MakeAskToChangeModelCli()
}

func init() {
	cmd.AddCommand(resetCmd)
	cmd.AddCommand(setCmd)
}
