package cmd

import (
	"fmt"

	"github.com/christian-gama/autocommit/internal/groq"
	"github.com/christian-gama/autocommit/internal/openai"
	"github.com/christian-gama/autocommit/internal/provider"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration file",
	Run:   runReset,
}

func runReset(cmd *cobra.Command, args []string) {
	llmProvider := loadProvider()
	proiderFactory := provider.NewProviderFactory(llmProvider)
	configRepo := proiderFactory.MakeConfigRepo()
	switch llmProvider.GetName() {
	case "openai":
		fmt.Println("⌛ Resetting OpenAI configuration file...")
		resetConfigCommand = openai.NewOpenAIResetConfigCommand(configRepo)
	case "groq":
		fmt.Println("⌛ Resetting Groq configuration file...")
		resetConfigCommand = groq.NewGroqResetConfigCommand(configRepo)
	}
	if err := resetConfigCommand.Execute(); err != nil {
		panic(err)
	}

	fmt.Println(
		"✅ Configuration file reset successfully - Next time you run autocommit, you will be asked to configure it again.",
	)
}
