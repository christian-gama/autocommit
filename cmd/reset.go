package cmd

import (
	"fmt"

	"github.com/christian-gama/autocommit/internal/groq"
	"github.com/christian-gama/autocommit/internal/llm"
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
	var rsg llm.ResetConfigCommand
	switch llmProvider.GetName() {
	case "openai":
		fmt.Println("⌛ Resetting OpenAI configuration file...")
		rsg = openai.NewOpenAIResetConfigCommand(configRepo)
	case "groq":
		fmt.Println("⌛ Resetting Groq configuration file...")
		rsg = groq.NewGroqResetConfigCommand(configRepo)
	}
	if err := rsg.Execute(); err != nil {
		panic(err)
	}

	fmt.Println(
		"✅ Configuration file reset successfully - Next time you run autocommit, you will be asked to configure it again.",
	)
}
