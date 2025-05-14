package provider

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/v2/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/mistral"
)

type Mistral struct{}

func (m Mistral) New(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(m.Name())
	if !ok {
		return nil, fmt.Errorf("no Mistral LLM provider found")
	}

	if err := os.Setenv("MISTRAL_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set MISTRAL_API_KEY: %w", err)
	}

	return mistral.New(mistral.WithModel(llm.Model))
}
func (m Mistral) Name() string {
	return "Mistral"
}

func (m Mistral) Models() []string {
	return []string{
		"mistral-large-latest",
		"mistral-medium-latest",
		"mistral-small-latest",
	}
}
