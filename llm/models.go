package llm

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

const _openai = "openai"

func makeOpenAI(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(_openai)
	if !ok {
		return nil, fmt.Errorf("no OpenAI LLM provider found")
	}

	if err := os.Setenv("OPENAI_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set OPENAI_API_KEY: %w", err)
	}

	return openai.New(openai.WithModel(llm.Model))
}
