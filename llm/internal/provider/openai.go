package provider

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/v2/config"
	"github.com/tmc/langchaingo/llms"
	openai "github.com/tmc/langchaingo/llms/openai"
)

type OpenAI struct{}

func (o OpenAI) New(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(o.Name())
	if !ok {
		return nil, fmt.Errorf("no OpenAI LLM provider found")
	}

	if err := os.Setenv("OPENAI_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set OPENAI_API_KEY: %w", err)
	}

	return openai.New(openai.WithModel(llm.Model))
}

func (o OpenAI) Name() string {
	return "OpenAI"
}

func (o OpenAI) Models() []string {
	return []string{
		"gpt-5",
		"gpt-5-mini",
		"gpt-5-nano",
		"gpt-5.2",
	}
}
