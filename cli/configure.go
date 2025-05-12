package cli

import (
	"github.com/christian-gama/autocommit/ask"
	"github.com/christian-gama/autocommit/config"
	"github.com/spf13/cobra"
)

var Configure = &cobra.Command{
	Use:   "configure",
	Short: "Configure an existing LLM provider or add a new one",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _, err := config.LoadOrNew()
		if err != nil {
			cmd.PrintErrf("Error loading config: %v\n", err)
			return
		}

		if err := configure(cfg); err != nil {
			cmd.PrintErrf("Error configuring LLM provider: %v\n", err)
			return
		}

		cmd.Println("LLM provider configured successfully.")
	},
}

func configure(cfg *config.Config) error {
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

	model, err := askConfig.Model(provider, defaults.Model)
	if err != nil {
		return err
	}

	credential, err := askConfig.Credential(defaults.Credential)
	if err != nil {
		return err
	}

	isDefault, err := askConfig.IsDefault(defaults.IsDefault)
	if err != nil {
		return err
	}

	if err := cfg.SetLLM(provider, model, credential, isDefault); err != nil {
		return err
	}

	if err := cfg.Save(); err != nil {
		return err
	}

	return nil
}
