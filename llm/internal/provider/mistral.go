package provider

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/mistral"
)

// Mistral is the identifier for the Mistral AI LLM provider.
const Mistral = "Mistral"

// MakeMistral creates and configures a Mistral language model instance.
func MakeMistral(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(Mistral)
	if !ok {
		return nil, fmt.Errorf("no Mistral LLM provider found")
	}

	if err := os.Setenv("MISTRAL_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set MISTRAL_API_KEY: %w", err)
	}

	return mistral.New(mistral.WithModel(llm.Model))
}

var MistralModels = []string{
	"mistral-large-latest",
	"mistral-medium-latest",
	"mistral-small-latest",
}
