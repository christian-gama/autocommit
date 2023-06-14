package cli

import (
	"log"

	"github.com/christian-gama/autocommit/chat"
	"github.com/christian-gama/autocommit/store"
	"github.com/spf13/cobra"
)

var set = &cobra.Command{
	Use:   "set",
	Short: "Set configuration options",
	Run: func(cmd *cobra.Command, args []string) {
		configStore := store.NewStore()

		if !configStore.IsStored() {
			log.Fatal("Config file not found - run 'autocommit' to set it up")
		}

		if OpenAIAPIKey != "" {
			if err := chat.ValidateAPIKey(OpenAIAPIKey); err != nil {
				log.Fatal(err)
			}

			configStore.SetOpenAIAPIKey(OpenAIAPIKey)
		}

		if OpenAIModel != "" {
			if err := chat.ValidateModel(OpenAIModel); err != nil {
				log.Fatal(err)
			}

			configStore.SetOpenAIModel(OpenAIModel)
		}

		if OpenAITemperature != 0.0 {
			if err := chat.ValidateTemperature(OpenAITemperature); err != nil {
				log.Fatal(err)
			}

			configStore.SetOpenAITemperature(OpenAITemperature)
		}
	},
}

var (
	OpenAIAPIKey      string
	OpenAIModel       string
	OpenAITemperature float32
)

func init() {
	set.Flags().StringVarP(&OpenAIAPIKey, "openai-api-key", "k", "", "openAI API Key")
	set.Flags().StringVarP(&OpenAIModel, "openai-model", "m", "", "openAI Model")
	set.Flags().Float32VarP(&OpenAITemperature, "openai-temperature", "t", 0.0, "openAI Temperature")
}
