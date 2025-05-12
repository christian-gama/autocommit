package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/christian-gama/autocommit/ask"
	"github.com/christian-gama/autocommit/autocommit"
	"github.com/christian-gama/autocommit/config"
	"github.com/christian-gama/autocommit/llm"
	"github.com/spf13/cobra"
)

var AutoCommit = &cobra.Command{
	Use:   "autocommit",
	Short: "Autocommit is a CLI tool that uses LLM models to generate commit messages based on the changes made in your current repository.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			if errors.Is(err, config.ErrConfigNotFound) {
				provider, err := ask.Provider("")
				if err != nil {
					panic(err)
				}

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

				cfg, err = config.New()
				if err != nil {
					panic(err)
				}

				if err := cfg.SetLLM(provider, model, credential, isDefault); err != nil {
					panic(err)
				}

				if err := cfg.Save(); err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		}

		model, err := llm.New(cfg)
		if err != nil {
			panic(err)
		}

		autocommit, err := autocommit.New(model)
		if err != nil {
			panic(err)
		}

		completion, err := autocommit.Generate(context.Background())
		if err != nil {
			panic(err)
		}

		fmt.Println(completion)
	},
}
