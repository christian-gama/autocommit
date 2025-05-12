package cli

import (
	"context"
	"fmt"

	"github.com/christian-gama/autocommit/config"
	"github.com/christian-gama/autocommit/generator"
	"github.com/christian-gama/autocommit/llm"
	"github.com/spf13/cobra"
)

var AutoCommit = &cobra.Command{
	Use:   "autocommit",
	Short: "Autocommit is a CLI tool that uses LLM models to generate commit messages based on the changes made in your current repository.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, isNew, err := config.LoadOrNew()
		if err != nil {
			cmd.PrintErrf("Error loading config: %v\n", err)
			return
		}

		if isNew {
			if err := configure(cfg); err != nil {
				cmd.PrintErrf("Error configuring LLM provider: %v\n", err)
				return
			}
		}

		model, err := llm.New(cfg)
		if err != nil {
			cmd.PrintErrf("Error creating LLM model: %v\n", err)
			return
		}

		generator, err := generator.New(model)
		if err != nil {
			cmd.PrintErrf("Error creating generator: %v\n", err)
		}

		completion, err := generator.Generate(context.Background())
		if err != nil {
			cmd.PrintErrf("Error generating commit message: %v\n", err)
			return
		}

		fmt.Println(completion)
	},
}
