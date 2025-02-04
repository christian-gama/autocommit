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
	OpenAIAPIKey string
	OpenAIModel  string
)

func init() {
	setCmd.Flags().
		StringVarP(&OpenAIAPIKey, "api-key", "k", "", "openAI API Key")
	setCmd.Flags().StringVarP(&OpenAIModel, "model", "m", "", "openAI Model")

	setCmd.RegisterFlagCompletionFunc(
		"model",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return openai.AllowedModels, cobra.ShellCompDirectiveNoFileComp
		},
	)
}

func runSet(cmd *cobra.Command, args []string) {
	config := openai.NewConfig(OpenAIAPIKey, OpenAIModel)

	if err := updateConfigCommand.Execute(config); err != nil {
		panic(err)
	}
}
