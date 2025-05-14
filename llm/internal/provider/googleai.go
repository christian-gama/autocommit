package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

// GoogleAI is the identifier for the Google AI LLM provider.
const GoogleAI = "Google AI"

// MakeGoogleAI creates and configures a Google AI language model instance.
func MakeGoogleAI(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(GoogleAI)
	if !ok {
		return nil, fmt.Errorf("no Google AI LLM provider found")
	}

	if err := os.Setenv("API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set API_KEY: %w", err)
	}

	return googleai.New(context.Background(), googleai.WithAPIKey(llm.Credential), googleai.WithDefaultModel(llm.Model))
}

var GoogleAIModels = []string{
	"mistral-large-latest",
	"mistral-medium-latest",
	"mistral-small-latest",
}
