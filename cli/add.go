package cli

import (
	"github.com/christian-gama/autocommit/ask"
	"github.com/christian-gama/autocommit/config"
	"github.com/spf13/cobra"
)

var Add = &cobra.Command{
	Use:   "add",
	Short: "Add a new LLM provider",
	Long:  "Add a new LLM provider to the configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			panic(err)
		}

		defaultLLM, err := cfg.DefaultLLM()
		if err != nil {
			panic(err)
		}

		provider, err := ask.Provider("")
		if err != nil {
			panic(err)
		}

		if provider == defaultLLM.Provider() {
			model, err := ask.Model(provider, defaultLLM.Model())
			if err != nil {
				panic(err)
			}

			credential, err := ask.Credential(defaultLLM.Credential())
			if err != nil {
				panic(err)
			}

			isDefault, err := ask.IsDefault(defaultLLM.IsDefault())
			if err != nil {
				panic(err)
			}

			if err := cfg.SetLLM(provider, model, credential, isDefault); err != nil {
				panic(err)
			}
		} else {
			model, err := ask.Model(provider, "")
			if err != nil {
				panic(err)
			}

			credential, err := ask.Credential("")
			if err != nil {
				panic(err)
			}

			isDefault, err := ask.IsDefault(false)
			if err != nil {
				panic(err)
			}

			if err := cfg.SetLLM(provider, model, credential, isDefault); err != nil {
				panic(err)
			}
		}

		if err := cfg.Save(); err != nil {
			panic(err)
		}
	},
}

func init() {
	AutoCommit.AddCommand(Add)
}
