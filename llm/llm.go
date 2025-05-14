// Package llm provides integration with various language model providers.
package llm

import (
	"fmt"

	"github.com/christian-gama/autocommit/config"
	"github.com/christian-gama/autocommit/llm/internal/provider"
	"github.com/tmc/langchaingo/llms"
)

// Providers is a registry of available LLM providers.
var Providers = newProviderRegistry([]provider.Provider{
	provider.OpenAI{},
	provider.GoogleAI{},
	provider.Mistral{},
	provider.Groq{},
	provider.Ollama2{},
})

type providerRegistry map[string]provider.Provider

func newProviderRegistry(providers []provider.Provider) providerRegistry {
	m := make(providerRegistry, len(providers))
	for _, provider := range providers {
		if _, exists := m[provider.Name()]; exists {
			panic(fmt.Sprintf("duplicate provider name: %s", provider.Name()))
		}

		m[provider.Name()] = provider
	}

	return m
}

// Models returns a list of available models for the specified provider.
func (r providerRegistry) Models(providerName string) []string {
	p, ok := r[providerName]
	if !ok {
		return nil
	}

	return p.Models()
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
func (r providerRegistry) New(cfg *config.Config) (llms.Model, error) {
	defaultLLM, ok := cfg.DefaultLLM()
	if !ok {
		return nil, fmt.Errorf("no default LLM provider found")
	}

	p, ok := r[defaultLLM.Provider]
	if !ok {
		return nil, fmt.Errorf("unsupported provider: %s", defaultLLM.Provider)
	}

	return p.New(cfg)
}
