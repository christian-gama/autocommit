package llm

import (
	"fmt"

	"github.com/christian-gama/autocommit/config"
	"github.com/tmc/langchaingo/llms"
)

func New(config *config.Config) (llms.Model, error) {
	defaultLLM, ok := config.DefaultLLM()
	if !ok {
		return nil, fmt.Errorf("no default LLM provider found")
	}

	switch defaultLLM.Provider {
	case OpenAI:
		return makeOpenAI(config)
	case Ollama2:
		return makeOllama(config)
	case Mistral:
		return makeMistral(config)
	case GoogleAI:
		return makeGoogleAI(config)
	default:
		return nil, fmt.Errorf("unsupported LLM provider: %s", defaultLLM.Provider)
	}
}
