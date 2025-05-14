package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

type GoogleAI struct{}

func (g GoogleAI) New(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(g.Name())
	if !ok {
		return nil, fmt.Errorf("no Google AI LLM provider found")
	}

	if err := os.Setenv("API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set API_KEY: %w", err)
	}

	return googleai.New(context.Background(), googleai.WithAPIKey(llm.Credential), googleai.WithDefaultModel(llm.Model))
}

func (g GoogleAI) Name() string {
	return "Google AI"
}

func (g GoogleAI) Models() []string {
	return []string{
		"gemini-2.0-flash",
		"gemini-2.5-pro-exp-03-25",
		"gemini-2.5-pro-preview-05-06",
		"gemini-2.5-flash-preview-04-17",
	}
}
