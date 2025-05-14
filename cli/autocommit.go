package cli

import (
	"context"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/christian-gama/autocommit/ask"
	"github.com/christian-gama/autocommit/config"
	"github.com/christian-gama/autocommit/generator"
	"github.com/christian-gama/autocommit/git"
	"github.com/christian-gama/autocommit/instruction"
	"github.com/christian-gama/autocommit/llm"
	"github.com/spf13/cobra"
)

// AutoCommitCmd is the main command for the autocommit tool. It uses LLM models to
// generate commit messages based on the changes made in the current Git repository.
var AutoCommitCmd = &cobra.Command{
	Use:                   "autocommit",
	Short:                 "Generate AI-powered git commit messages interactively",
	Long:                  description,
	Example:               example,
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cobra.NoFileCompletions,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !config.HasConfig() {
			if err := runConfigure(); err != nil {
				return err
			}
		}

		if !instruction.HasInstruction() {
			if err := runRestoreInstructions(); err != nil {
				return err
			}
		}

		currentModel, ok := _config.CurrentModel()
		if !ok {
			return fmt.Errorf("error getting default LLM model")
		}

		fmt.Printf("🤖 Using model: %s\n", currentModel)

		model, err := llm.Providers.New(_config)
		if err != nil {
			return fmt.Errorf("error creating LLM model: %w", err)
		}

		generator, err := generator.New(model)
		if err != nil {
			return fmt.Errorf("error creating generator: %w", err)
		}

		_generator = generator
		_askAction = ask.NewAction()

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		completion, err := handleGeneration()
		if err != nil {
			return err
		}

		return handleInteraction(completion)
	},
}

func handleInteraction(completion string) (err error) {
	for {
		action, err := _askAction.Action()
		if err != nil {
			return err
		}

		switch action {
		case ask.ActionAddInstruction:
			{
				completion, err = handleAddInstruction()
				if err != nil {
					return err
				}
				continue
			}

		case ask.ActionCommit:
			{
				return handleCommit(completion)
			}

		case ask.ActionCopyToClipboard:
			{
				return handleCopyToClipboard(completion)
			}

		case ask.ActionRegenerate:
			{
				completion, err = handleGeneration("Regenerate the commit message with a different output.")
				if err != nil {
					return err
				}
				continue
			}

		case ask.ActionExit:
			{
				return nil
			}

		default:
			panic(fmt.Sprintf("unexpected ask.ActionOption: %#v", action))
		}

	}
}

func handleGeneration(additionalInstructions ...string) (string, error) {
	completion, err := _generator.Generate(context.Background(), additionalInstructions...)
	if err != nil {
		return "", err
	}

	fmt.Printf("💬 Commit message:"+
		"\n==================================================================================================\n%s"+
		"\n==================================================================================================\n",
		completion,
	)

	return completion, nil
}

func handleAddInstruction() (string, error) {
	instruction, err := _askAction.Instruction()
	if err != nil {
		return "", err
	}

	return handleGeneration(instruction)
}

func handleCommit(completion string) error {
	return git.Commit(completion)
}

func handleCopyToClipboard(completion string) error {
	if err := clipboard.WriteAll(completion); err != nil {
		return err
	}

	fmt.Println("✅ Commit message copied to clipboard.")

	return nil
}
