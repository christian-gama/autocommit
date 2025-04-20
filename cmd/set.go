package cmd

import (
	"log"

	"github.com/christian-gama/autocommit/internal/groq"
	"github.com/christian-gama/autocommit/internal/openai"
	"github.com/christian-gama/autocommit/internal/provider"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration configs",
	Run:   runSet,
}

var (
	llmAPIKey      string
	llmModel       string
	llmTemperature float32
	providerName   string
)

func init() {
	setCmd.Flags().
		StringVarP(&providerName, "provider", "p", "", "LLM provider (e.g., openai, groq)")
	setCmd.Flags().
		StringVarP(&llmAPIKey, "apikey", "k", "", "API key for the LLM provider")
	setCmd.Flags().StringVarP(&llmModel, "model", "m", "", "Model to use for the LLM provider")
	setCmd.Flags().
		Float32VarP(&llmTemperature, "temperature", "t", 0.7, "Temperature for the LLM provider")

	if err := setCmd.RegisterFlagCompletionFunc(
		"model",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if llmProvider == nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return llmProvider.GetAllowedModels(), cobra.ShellCompDirectiveNoFileComp
		},
	); err != nil {
		log.Fatalf("Failed to register flag completion: %v", err)
	}
}

func runSet(cmd *cobra.Command, args []string) {
	if providerName == "" || llmAPIKey == "" || llmModel == "" || llmTemperature == 0 {
		log.Fatal("Missing one or more required flags")
	}

	switch providerName {
	case "openai":
		llmProvider = openai.NewOpenAIProvider()
		updateConfigCommand = openai.NewOpenAIUpdateConfigCommand(
			provider.NewProviderFactory(llmProvider).MakeConfigRepo(),
		)
	case "groq":
		llmProvider = groq.NewGroqProvider()
		updateConfigCommand = groq.NewGroqUpdateConfigCommand(
			provider.NewProviderFactory(llmProvider).MakeConfigRepo(),
		)

	default:
		log.Fatalf("Unsupported provider: %s", providerName)
	}

	if llmProvider == nil {
		log.Fatal("Failed to initialize LLM provider")
	}

	config := llmProvider.NewConfig(llmAPIKey, llmModel, llmTemperature)
	if config == nil {
		log.Fatal("Failed to create configuration")
	}

	if err := updateConfigCommand.Execute(config); err != nil {
		log.Fatalf("Failed to execute updateConfigCommand: %v", err)
	}
}
