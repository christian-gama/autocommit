package provider

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/config"
	"github.com/tmc/langchaingo/llms"
	openai "github.com/tmc/langchaingo/llms/openai"
)

// Groq is the identifier for the Groq LLM provider.
const Groq = "Groq"

// MakeGroq creates and configures an Groq language model instance.
func MakeGroq(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(Groq)
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

var GroqModels = []string{
	"gemma2-9b-it",
	"llama-3.3-70b-versatile",
	"llama-3.1-8b-instant",
	"llama3-70b-8192",
	"llama3-8b-8192",
}
