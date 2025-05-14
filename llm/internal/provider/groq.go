package provider

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/v2/config"
	"github.com/tmc/langchaingo/llms"
	openai "github.com/tmc/langchaingo/llms/openai"
)

type Groq struct{}

func (g Groq) New(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(g.Name())
	if !ok {
		return nil, fmt.Errorf("no Groq LLM provider found")
	}

	if err := os.Setenv("GROQ_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set GROQ_API_KEY: %w", err)
	}

	return openai.New(
		openai.WithModel(llm.Model),
		openai.WithToken(llm.Credential),
		openai.WithBaseURL("https://api.groq.com/openai/v1"),
	)
}

func (Groq) Name() string {
	return "Groq"
}

func (Groq) Models() []string {
	return []string{
		"gemma2-9b-it",
		"llama-3.3-70b-versatile",
		"llama-3.1-8b-instant",
		"llama3-70b-8192",
		"llama3-8b-8192",
	}
}
