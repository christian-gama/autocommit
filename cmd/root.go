package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/christian-gama/autocommit/internal/autocommit"
	"github.com/christian-gama/autocommit/internal/openai"
	"github.com/christian-gama/autocommit/internal/provider"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "autocommit",
	Run:   runCmd,
	Short: "Autocommit is a CLI tool that uses OpenAI's models to generate commit messages based on the changes made in the repository.",
}

// Execute executes the root command.
func Execute() error {
	return cmd.Execute()
}

func runCmd(cmd *cobra.Command, args []string) {
	llmProvider := loadProvider()
	providerFactory := provider.NewProviderFactory(llmProvider)
	generatorCommand = autocommit.MakeGeneratorCommand(llmProvider)
	verifyConfigCommand := providerFactory.MakeVerifyConfigCommand()
	askConfigsCli := providerFactory.MakeAskConfigsCli()
	systemMsgHealthCheck := autocommit.MakeSystemMsgHealthCheckCommand()

	var err error
	config, err = verifyConfigCommand.Execute(askConfigsCli.Execute)
	if err != nil {
		panic(err)
	}

	if err := systemMsgHealthCheck.Execute(); err != nil {
		panic(err)
	}

	fmt.Printf("🤖 Using model: %s\n", config.GetModel())

	handleCmd(cmd, args)
}

func handleCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("⌛ Creating a commit message...\n")
	response, err := generatorCommand.Execute(config)
	if err != nil {
		handleMaxToken(err, cmd, args)
	}

	printSuccessMessage(response)
	handlePostCommit(response)
}

func handlePostCommit(response string) {
	option, err := postCommitCli.Execute()
	if err != nil {
		panic(err)
	}

	switch option {
	case autocommit.CommitChangesOption:
		handleCommit(response)

	case autocommit.CopyToClipboardOption:
		handleCopyToClipboard(response)

	case autocommit.ExitOption:
		os.Exit(0)
	}
}

func handleCommit(response string) {
	if err := gitCommand.Commit(response); err != nil {
		panic(err)
	}
}

func handleCopyToClipboard(response string) {
	if err := clipboardCommand.Execute(response); err != nil {
		panic(err)
	}
}

func handleMaxToken(err error, cmd *cobra.Command, args []string) {
	isTokenError := strings.Contains(
		err.Error(),
		"Please reduce the length of the messages",
	)

	if !isTokenError || slices.Contains(openai.AllowedModels, config.GetModel()) {
		panic(err)
	}

	answer, err := askToChangeModelCli.Execute()
	if err != nil {
		panic(err)
	}

	if answer {
		handleCmd(cmd, args)
		return
	}

	panic(
		"🚧 You reached the maximum allowed token for this model. You can try a new model running 'autocommit set --model <model>' or decrease the amount of files being commited.",
	)
}

func printSuccessMessage(response string) {
	fmt.Printf("📝 Commit message generated: \n%s\n\n", response)
}
