package provider

import (
	"github.com/christian-gama/autocommit/config"
	"github.com/tmc/langchaingo/llms"
)

// Provider defines the contract for an LLM provider: its unique name,
// available models, and factory method to create a configured llms.Model.
type Provider interface {
	// Name returns the unique identifier for this provider.
	Name() string
	// Models returns the list of supported model identifiers.
	Models() []string
	// New creates and configures a new llms.Model using the given config.
	New(*config.Config) (llms.Model, error)
}
