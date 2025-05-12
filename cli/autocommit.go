package cli

import (
	"context"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/christian-gama/autocommit/ask"
	"github.com/christian-gama/autocommit/config"
	"github.com/christian-gama/autocommit/generator"
	"github.com/christian-gama/autocommit/git"
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
			return
		}

		completion, err := generator.Generate(context.Background())
		if err != nil {
			cmd.PrintErrf("Error generating commit message: %v\n", err)
			return
		}

		cmd.Printf("Generated commit message:\n%s", completion)

		askAction := ask.NewAction()

		for {
			action, err := askAction.Action()
			if err != nil {
				cmd.PrintErrf("Error asking for action: %v\n", err)
				return
			}

			switch action {
			case ask.ActionAddInstruction:
				instruction, err := askAction.Instruction()
				if err != nil {
					cmd.PrintErrf("Error asking for instruction: %v\n", err)
					return
				}

				completion, err = generator.Generate(context.Background(), instruction)
				if err != nil {
					cmd.PrintErrf("Error generating commit message: %v\n", err)
					return
				}

				cmd.Println("Generated commit message:\n", completion)
			case ask.ActionCommit:
				if err := git.Commit(completion); err != nil {
					cmd.PrintErrf("Error committing changes: %v\n", err)
					return
				}

				return
			case ask.ActionCopyToClipboard:
				err := clipboard.WriteAll(fmt.Sprintf("git commit -m %q", completion))
				if err != nil {
					cmd.PrintErrf("Error copying to clipboard: %v\n", err)
					return
				}

				return
			case ask.ActionRegenerate:
				completion, err := generator.Generate(context.Background(), "Regenerate the commit message with a different output.")
				if err != nil {
					cmd.PrintErrf("Error generating commit message: %v\n", err)
					return
				}

				cmd.Println("Generated commit message:\n", completion)
			case ask.ActionExit:
				return
			default:
				panic(fmt.Sprintf("unexpected ask.ActionOption: %#v", action))
			}
		}
	},
}
