package llm

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func New(config *config.Config) (llms.Model, error) {
	active, err := config.DefaultLLM()
	if err != nil {
		return nil, err
	}

	switch active.Provider() {
	case "openai":
		return makeOpenAI(config)
	default:
		return nil, fmt.Errorf("unsupported LLM provider: %s", active.Provider())
	}
}

func makeOpenAI(config *config.Config) (llms.Model, error) {
	llm, err := config.LLM("openai")
	if err != nil {
		return nil, err
	}

	if err := os.Setenv("OPENAI_API_KEY", llm.Credential()); err != nil {
		return nil, fmt.Errorf("set OPENAI_API_KEY: %w", err)
	}

	return openai.New(openai.WithModel(llm.Model()))
}
