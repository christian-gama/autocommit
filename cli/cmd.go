package cli

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/christian-gama/autocommit/chat"
	"github.com/christian-gama/autocommit/git"
	"github.com/christian-gama/autocommit/store"
	"github.com/spf13/cobra"
)

var verbose bool
var cmd = &cobra.Command{
	Use:   "",
	Run:   run,
	Short: "Autocommit is a CLI tool that uses OpenAI's models to generate commit messages based on the changes made in the repository.",
}

func init() {
	cmd.AddCommand(reset)
	cmd.AddCommand(set)
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "create a verbose commit message")
}

// Execute executes the root command.
func Execute() error {
	return cmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	config := loadConfig()
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

	chatAnswer := askUserForChatOption()
	handleChatOption(cmd, chatAnswer.Option, commitMessage)
}

func loadConfig() *store.Config {
	configStore := store.NewStore()

	if !configStore.IsStored() {
		configAnswers := askUserForConfig()
		configStore.CreateConfigFile(&store.Config{
			OpenAIAPIKey:      configAnswers.OpenAIAPIKey,
			OpenAIModel:       configAnswers.Model,
			OpenAITemperature: configAnswers.Temperature,
		})
	}

	return configStore.Config()
}

func handleChatOption(cmd *cobra.Command, option, commitMessage string) {
	switch option {
	case commitChangesOption:
		git.Commit(commitMessage)

	case generateCommitMessageOption:
		run(cmd, nil)

	case copyCommitMessageToClipboardOption:
		if err := clipboard.WriteAll(fmt.Sprintf("git commit -m \"%s\"", commitMessage)); err != nil {
			log.Fatalf("Failed to copy commit message to clipboard: %v", err)
		}
	}
}
