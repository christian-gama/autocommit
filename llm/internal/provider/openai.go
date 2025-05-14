package provider

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/config"
	goopenai "github.com/sashabaranov/go-openai"
	"github.com/tmc/langchaingo/llms"
	openai "github.com/tmc/langchaingo/llms/openai"
)

// OpenAI is the identifier for the OpenAI LLM provider.
const OpenAI = "OpenAI"

// MakeOpenAI creates and configures an OpenAI language model instance.
func MakeOpenAI(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(OpenAI)
	if !ok {
		return nil, fmt.Errorf("no OpenAI LLM provider found")
	}

	if err := os.Setenv("OPENAI_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set OPENAI_API_KEY: %w", err)
	}

	return openai.New(openai.WithModel(llm.Model))
}

var OpenAIModels = []string{
	goopenai.GPT4o,
	goopenai.GPT4Dot1,
	goopenai.GPT4Dot1Mini,
	goopenai.GPT4Dot1Nano,
	goopenai.O1,
	goopenai.O1Mini,
	goopenai.O3,
	goopenai.O3Mini,
	goopenai.O4Mini,
}
