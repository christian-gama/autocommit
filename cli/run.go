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
		if err := clipboard.WriteAll(commitMessage); err != nil {
			log.Fatalf("Failed to copy commit message to clipboard: %v", err)
		}
	}
}
