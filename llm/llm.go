// Package llm provides integration with various language model providers.
package llm

import (
	"fmt"

	"github.com/christian-gama/autocommit/config"
	"github.com/christian-gama/autocommit/llm/internal/provider"
	"github.com/tmc/langchaingo/llms"
)

type providerRegistry map[string]struct {
	factory provider.Func
	models  []string
}

// Models returns a list of available models for the specified provider.
func (r providerRegistry) Models(provider string) []string {
	models, ok := r[provider]
	if !ok {
		return nil
	}

	return models.models
}

// List returns a list of all available LLM providers.
func (r providerRegistry) List() []string {
	providers := make([]string, 0, len(r))
	for provider := range r {
		providers = append(providers, provider)
	}
	return providers
}

// New creates a new LLM model based on the default provider specified in the
func (r providerRegistry) New(config *config.Config) (llms.Model, error) {
	defaultLLM, ok := config.DefaultLLM()
	if !ok {
		return nil, fmt.Errorf("no default LLM provider found")
	}

	providerConfig, ok := r[defaultLLM.Provider]
	if !ok {
		return nil, fmt.Errorf("unsupported provider: %s", defaultLLM.Provider)
	}

	return providerConfig.factory(config)
}

// Providers is a registry of available LLM providers and their corresponding
var Providers = providerRegistry{
	provider.OpenAI: {
		factory: provider.MakeOpenAI,
		models:  provider.OpenAIModels,
	},

	provider.GoogleAI: {
		factory: provider.MakeGoogleAI,
		models:  provider.GoogleAIModels,
	},

	provider.Ollama2: {
		factory: provider.MakeOllama2,
		models:  provider.Ollama2Models,
	},

	provider.Mistral: {
		factory: provider.MakeMistral,
		models:  provider.MistralModels,
	},
}
