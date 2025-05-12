package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/config"
	"github.com/sashabaranov/go-openai"
	"github.com/tmc/langchaingo/llms"
	lcgoogleai "github.com/tmc/langchaingo/llms/googleai"
	lcmistral "github.com/tmc/langchaingo/llms/mistral"
	lcollama "github.com/tmc/langchaingo/llms/ollama"
	lcopenai "github.com/tmc/langchaingo/llms/openai"
)

const OpenAI = "OpenAI"

func makeOpenAI(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(OpenAI)
	if !ok {
		return nil, fmt.Errorf("no OpenAI LLM provider found")
	}

	if err := os.Setenv("OPENAI_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set OPENAI_API_KEY: %w", err)
	}

	return lcopenai.New(lcopenai.WithModel(llm.Model))
}

const Ollama2 = "Ollama 2"

func makeOllama(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(Ollama2)
	if !ok {
		return nil, fmt.Errorf("no Ollama LLM provider found")
	}

	if err := os.Setenv("OLLAMA_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set OLLAMA_API_KEY: %w", err)
	}

	return lcollama.New(lcollama.WithModel(llm.Model))
}

const Mistral = "Mistral"

func makeMistral(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(Mistral)
	if !ok {
		return nil, fmt.Errorf("no Mistral LLM provider found")
	}

	if err := os.Setenv("MISTRAL_API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set MISTRAL_API_KEY: %w", err)
	}

	return lcmistral.New(lcmistral.WithModel(llm.Model))
}

const GoogleAI = "Google AI"

func makeGoogleAI(config *config.Config) (llms.Model, error) {
	llm, ok := config.LLM(GoogleAI)
	if !ok {
		return nil, fmt.Errorf("no Google AI LLM provider found")
	}

	if err := os.Setenv("API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set API_KEY: %w", err)
	}

	return lcgoogleai.New(context.Background(), lcgoogleai.WithAPIKey(llm.Credential), lcgoogleai.WithDefaultModel(llm.Model))
}

func Models(provider string) []string {
	switch provider {
	case OpenAI:
		return []string{
			openai.GPT4o,
			openai.GPT4Dot1,
			openai.GPT4Dot1Mini,
			openai.GPT4Dot1Nano,
			openai.O1,
			openai.O1Mini,
			openai.O3,
			openai.O3Mini,
			openai.O4Mini,
		}
	case Ollama2:
		return []string{
			"gemma:1b",
			"gemma:4b",
			"gemma:12b",
			"gemma:27b",
			"qwen3:0.6b",
			"qwen3:1.7b",
			"qwen3:4b",
			"qwen3:8b",
			"qwen3:14b",
			"qwen3:30b",
			"qwen3:32b",
			"qwen3:235b",
			"deepseek-r1:1.5b",
			"deepseek-r1:7b",
			"deepseek-r1:8b",
			"deepseek-r1:14b",
			"deepseek-r1:32b",
			"deepseek-r1:70b",
			"deepseek-r1:671b",
			"llama4",
			"llama3.3",
		}
	case Mistral:
		return []string{
			"mistral-large-latest",
			"mistral-medium-latest",
			"mistral-small-latest",
		}
	case GoogleAI:
		return []string{
			"gemini-2.0-flash",
			"gemini-2.5-pro-exp-03-25",
			"gemini-2.5-pro-preview-05-06",
			"gemini-2.5-flash-preview-04-17",
		}
	default:
		panic(fmt.Sprintf("unsupported provider: %s", provider))
	}
}
