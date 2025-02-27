package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/christian-gama/autocommit/internal/autocommit"
	"github.com/christian-gama/autocommit/internal/llm/openai"
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
	var err error
	config, err = verifyConfigCommand.Execute(askConfigsCli.Execute)
	if err != nil {
		panic(err)
	}

	if err := systemMsgHealthCheck.Execute(); err != nil {
		panic(err)
	}

	// Use type assertion to get provider-specific information for display
	if openAIConfig, ok := config.(*openai.Config); ok {
		fmt.Printf("🤖 Using model: %s\n", openAIConfig.Model)
	} else {
		fmt.Printf("🤖 Using %s provider\n", config.Provider())
	}

	handleCmd(cmd, args)
}

func handleCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("⌛ Creating a commit message...\n")
	response, err := generatorCommand.Execute(config)
	if err != nil {
		handleMaxToken(err, cmd, args)
	}

	printSuccessMessage(response)
	handlePostCommit(response, cmd, args)
}

func handlePostCommit(response string, cmd *cobra.Command, args []string) {
	option, err := postCommitCli.Execute()
	if err != nil {
		panic(err)
	}

	switch option {
	case autocommit.CommitChangesOption:
		handleCommit(response)

	case autocommit.CopyToClipboardOption:
		handleCopyToClipboard(response)

	case autocommit.RegenerateOption:
		handleRegenerate(cmd, args)

	case autocommit.AddInstructionOption:
		handleNewInstructions(cmd, args)

	case autocommit.ExitOption:
		os.Exit(0)
	}
}

func handleCommit(response string) {
	if err := commitCommand.Execute(response); err != nil {
		panic(err)
	}
}

func handleCopyToClipboard(response string) {
	if err := clipboardCommand.Execute(response); err != nil {
		panic(err)
	}
}

func handleNewInstructions(cmd *cobra.Command, args []string) {
	instructions, err := addInstructionCli.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("💡 Enhancing the message with your new instruction...\n")
	response, err := addInstructionCommand.Execute(config, instructions)
	if err != nil {
		panic(err)
	}

	printSuccessMessage(response)
	handlePostCommit(response, cmd, args)
}

func handleRegenerate(cmd *cobra.Command, args []string) {
	fmt.Printf("🔄 Regenerating the commit message...\n")
	response, err := addInstructionCommand.Execute(
		config,
		"Recreate the commit message from scratch. As a reminder, stick to the previous rules.",
	)
	if err != nil {
		panic(err)
	}

	printSuccessMessage(response)
	handlePostCommit(response, cmd, args)
}

func handleMaxToken(err error, cmd *cobra.Command, args []string) {
	// This is OpenAI-specific error handling
	// In a more generic solution, we might want to make this provider-specific too
	isTokenError := strings.Contains(
		err.Error(),
		"Please reduce the length of the messages",
	)

	openAIConfig, isOpenAI := config.(*openai.Config)
	if !isTokenError || !isOpenAI ||
		slices.Contains(openai.AllowedModels, openAIConfig.Model) {
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
