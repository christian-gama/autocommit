package provider

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/v2/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type Ollama2 struct{}

func (o Ollama2) New(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(o.Name())
	if !ok {
		return nil, fmt.Errorf("no Ollama LLM provider found")
	}

	if err := os.Setenv("OLLAMA_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set OLLAMA_API_KEY: %w", err)
	}

	return ollama.New(ollama.WithModel(llm.Model))
}

func (o Ollama2) Name() string {
	return "Ollama2"
}

func (o Ollama2) Models() []string {
	return []string{
		"gemma3",
		"qwen3",
		"deepseek-r1",
		"llama4",
		"llama3.3",
		"mistral",
	}
}
