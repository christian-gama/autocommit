// Package cli provides the command-line interface for the autocommit tool.
package cli

import (
	"fmt"

	"github.com/christian-gama/autocommit/ask"
	"github.com/christian-gama/autocommit/config"
	"github.com/spf13/cobra"
)

// Configure is a command that allows users to configure LLM providers.
// It guides the user through setting up or modifying provider details.
var Configure = &cobra.Command{
	Use:                   "configure",
	Short:                 "Configure an existing LLM provider or add a new one",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig()
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		if err := configureLLM(cfg); err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.Println("✅ LLM provider configured successfully.")
	},
	ValidArgsFunction: cobra.NoFileCompletions,
}

// loadConfig loads the existing configuration or creates a new one if none exists.
func loadConfig() (*config.Config, error) {
	cfg, _, err := config.LoadOrNew()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	return cfg, nil
}

// configureLLM guides the user through the LLM provider configuration process,
// prompting for provider, credential, model, and default status.
func configureLLM(cfg *config.Config) error {
	askConfig := ask.NewConfig()

	provider, err := askConfig.Provider()
	if err != nil {
		return err
	}

	var defaults struct {
		Model      string
		Credential string
		IsDefault  bool
	}

	if llm, ok := cfg.LLM(provider); ok {
		defaults.Model = llm.Model
		defaults.Credential = llm.Credential
		defaults.IsDefault = llm.IsDefault
	}

	credential, err := askConfig.Credential(defaults.Credential)
	if err != nil {
		return err
	}

	model, err := askConfig.Model(provider, defaults.Model)
	if err != nil {
		return err
	}

	var isDefault bool
	if cfg.HasAnyLLM() {
		isDefault, err = askConfig.IsDefault(defaults.IsDefault)
		if err != nil {
			return err
		}
	} else {
		// If there are no existing LLMs, set the first one as default
		// to avoid having to ask the user.
		isDefault = true
	}

	if err := cfg.SetLLM(provider, model, credential, isDefault); err != nil {
		return err
	}

	if err := cfg.Save(); err != nil {
		return err
	}

	return nil
}
