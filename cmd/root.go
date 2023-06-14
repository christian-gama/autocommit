package cmd

import (
	"fmt"
	"log"

	"github.com/christian-gama/autocommit/chat"
	"github.com/christian-gama/autocommit/config"
	"github.com/christian-gama/autocommit/git"
	"github.com/spf13/cobra"
)

var verbose bool
var cmd = &cobra.Command{
	Use:   "",
	Run:   run,
	Short: "Autocommit is a CLI tool that uses OpenAI's models to generate commit messages based on the changes made in the repository.",
}

func init() {
	cmd.AddCommand(resetCmd)
	cmd.AddCommand(setCmd)
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "create a verbose commit message")
}

// Execute executes the root command.
func Execute() error {
	return cmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	config := config.Load()
	chatService := chat.NewChatService(
		chat.NewConfig(config.OpenAIAPIKey, config.OpenAIModel, verbose, config.OpenAITemperature),
	)

	diff, err := git.Diff()
	if err != nil {
		log.Fatal(err)
	}

	commitMessage, err := chatService.Response(diff)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Commit message:\n'%s'\n\n", commitMessage)

	chatAnswer := chat.AskUserForChatOption()
	chat.HandleChatOption(func() { run(cmd, args) }, chatAnswer.Option, commitMessage)
}
