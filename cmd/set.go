package cmd

import (
	"github.com/christian-gama/autocommit/internal/openai"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration configs",
	Run:   runSet,
}

var (
	OpenAIAPIKey      string
	OpenAIModel       string
	OpenAITemperature float32
)

func init() {
	setCmd.Flags().StringVarP(&OpenAIAPIKey, "api-key", "k", "", "openAI API Key")
	setCmd.Flags().StringVarP(&OpenAIModel, "model", "m", "", "openAI Model")
	setCmd.Flags().Float32VarP(&OpenAITemperature, "temperature", "t", 0.0, "openAI Temperature")
}

func runSet(cmd *cobra.Command, args []string) {
	config := openai.NewConfig(OpenAIAPIKey, OpenAIModel, OpenAITemperature)

	if err := updateConfigCommand.Execute(config); err != nil {
		panic(err)
	}

}
