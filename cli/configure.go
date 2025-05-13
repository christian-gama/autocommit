// Package cli provides the command-line interface for the autocommit tool.
package cli

import (
	"github.com/spf13/cobra"
)

// configureCmd is a command that allows users to configureCmd LLM providers.
// It guides the user through setting up or modifying provider details.
var configureCmd = &cobra.Command{
	Use:                   "configure",
	Short:                 "Configure an existing LLM provider or add a new one",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runConfigure()
	},
}

func runConfigure() error {
	provider, err := _askConfig.Provider()
	if err != nil {
		return err
	}

	var defaults struct {
		Model      string
		Credential string
		IsDefault  bool
	}

	if llm, ok := _config.LLM(provider); ok {
		defaults.Model = llm.Model
		defaults.Credential = llm.Credential
		defaults.IsDefault = llm.IsDefault
	}

	credential, err := _askConfig.Credential(defaults.Credential)
	if err != nil {
		return err
	}

	model, err := _askConfig.Model(provider, defaults.Model)
	if err != nil {
		return err
	}

	var isDefault bool
	if _config.HasAnyLLM() {
		isDefault, err = _askConfig.IsDefault(defaults.IsDefault)
		if err != nil {
			return err
		}
	} else {
		// If there are no existing LLMs, set the first one as default
		// to avoid having to ask the user.
		isDefault = true
	}

	if err := _config.SetLLM(provider, model, credential, isDefault); err != nil {
		return err
	}

	if err := _config.Save(); err != nil {
		return err
	}

	return nil
}
