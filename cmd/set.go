package cmd

import (
	"fmt"
	"log"

	"github.com/christian-gama/autocommit/internal/groq"
	"github.com/christian-gama/autocommit/internal/openai"
	"github.com/christian-gama/autocommit/internal/provider"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [provider]",
	Short: "Set configuration options for a specific LLM provider",
	Long: `Set configuration options for a specific LLM provider.
The provider argument is required. Other flags are optional and only the specified ones will be updated.

Example usage:
  autocommit set openai --apikey your_api_key
  autocommit set groq --model llama3-70b-8192
  autocommit set openai --temperature 0.8`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("requires exactly one provider argument (openai, groq)")
		}

		validProviders := map[string]bool{
			"openai": true,
			"groq":   true,
		}

		if !validProviders[args[0]] {
			return fmt.Errorf("invalid provider: %s (must be one of: openai, groq)", args[0])
		}

		return nil
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return []string{"openai", "groq"}, cobra.ShellCompDirectiveNoFileComp
	},
	Run: runSet,
}

var (
	llmAPIKey      string
	llmModel       string
	llmTemperature float32
)

func init() {
	setCmd.Flags().
		StringVarP(&llmAPIKey, "apikey", "k", "", "API key for the LLM provider (optional)")

	setCmd.Flags().
		StringVarP(&llmModel, "model", "m", "", "Model to use for the LLM provider (optional)")

	setCmd.Flags().
		Float32VarP(&llmTemperature, "temperature", "t", 0.7, "Temperature for the LLM provider (optional, 0-1 range)")

	if err := setCmd.RegisterFlagCompletionFunc(
		"model",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}

			switch args[0] {
			case "openai":
				return openai.AllowedModels, cobra.ShellCompDirectiveNoFileComp

			case "groq":
				return groq.AllowedModels, cobra.ShellCompDirectiveNoFileComp

			default:
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
		},
	); err != nil {
		log.Fatalf("Failed to register flag completion: %v", err)
	}
}

func runSet(cmd *cobra.Command, args []string) {
	providerName := args[0]

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

	configToUpdate := llmProvider.NewConfig("", "", 0)

	flagsSet := false

	if cmd.Flags().Changed("apikey") {
		configToUpdate.SetAPIKey(llmAPIKey)
		flagsSet = true
	}

	if cmd.Flags().Changed("model") {
		configToUpdate.SetModel(llmModel)
		flagsSet = true
	}

	if cmd.Flags().Changed("temperature") {
		configToUpdate.SetTemperature(llmTemperature)
		flagsSet = true
	}

	if !flagsSet {
		log.Fatal("At least one configuration flag (--apikey, --model, or --temperature) must be specified")
	}

	if err := updateConfigCommand.Execute(configToUpdate); err != nil {
		log.Fatalf("Failed to execute updateConfigCommand: %v", err)
	}
}
