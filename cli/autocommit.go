// Package cli provides the command-line interface for the autocommit tool.
package cli

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/atotto/clipboard"
	"github.com/christian-gama/autocommit/ask"
	"github.com/christian-gama/autocommit/config"
	"github.com/christian-gama/autocommit/generator"
	"github.com/christian-gama/autocommit/git"
	"github.com/christian-gama/autocommit/llm"
	"github.com/spf13/cobra"
)

// AutoCommit is the main command for the autocommit tool. It uses LLM models to
// generate commit messages based on the changes made in the current Git repository.
var AutoCommit = &cobra.Command{
	Use:                   "autocommit",
	Short:                 "Autocommit is a CLI tool that uses LLM models to generate commit messages based on the changes made in your current repository.",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cobra.NoFileCompletions,
	Run: func(cmd *cobra.Command, args []string) {
		clearScreen()

		deps, err := newAutoCommitDeps(cmd)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		runAutoCommit(deps)
	},
}

// autoCommitDeps contains all dependencies required by the autocommit command.
type autoCommitDeps struct {
	cfg       *config.Config
	generator *generator.Generator
	askAction *ask.Action
	cmd       *cobra.Command
}

func newAutoCommitDeps(cmd *cobra.Command) (*autoCommitDeps, error) {
	cfg, isNew, err := config.LoadOrNew()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	if isNew {
		if err := configureLLM(cfg); err != nil {
			return nil, fmt.Errorf("error configuring LLM provider: %w", err)
		}
	}

	currentModel, ok := cfg.DefaultLLM()
	if !ok {
		return nil, fmt.Errorf("error getting default LLM model")
	}

	fmt.Printf("🤖 Using model: %s\n", currentModel.Model)

	model, err := llm.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating LLM model: %w", err)
	}

	generator, err := generator.New(model)
	if err != nil {
		return nil, fmt.Errorf("error creating generator: %w", err)
	}

	return &autoCommitDeps{
		cfg:       cfg,
		generator: generator,
		askAction: ask.NewAction(),
		cmd:       cmd,
	}, nil
}

func runAutoCommit(deps *autoCommitDeps) {
	completion, err := deps.generator.Generate(context.Background())
	if err != nil {
		deps.cmd.PrintErrf("Error generating commit message: %v\n", err)
		return
	}

	printCommitMessage(deps.cmd, completion)

	for {
		action, err := deps.askAction.Action()
		if err != nil {
			deps.cmd.PrintErrln(err)
			return
		}

		switch action {
		case ask.ActionAddInstruction:
			instruction, err := deps.askAction.Instruction()
			if err != nil {
				deps.cmd.PrintErrln(err)
				return
			}

			completion, err = deps.generator.Generate(context.Background(), instruction)
			if err != nil {
				deps.cmd.PrintErrln(err)
				return
			}

			printCommitMessage(deps.cmd, completion)
		case ask.ActionCommit:
			if err := git.Commit(completion); err != nil {
				deps.cmd.PrintErrln(err)
				return
			}

			return
		case ask.ActionCopyToClipboard:
			err := clipboard.WriteAll(fmt.Sprintf("git commit -m %q", completion))
			if err != nil {
				deps.cmd.PrintErrln(err)
				return
			}

			return
		case ask.ActionRegenerate:
			completion, err := deps.generator.Generate(context.Background(), "Regenerate the commit message with a different output.")
			if err != nil {
				deps.cmd.PrintErrln(err)
				return
			}

			printCommitMessage(deps.cmd, completion)
		case ask.ActionExit:
			return
		default:
			panic(fmt.Sprintf("unexpected ask.ActionOption: %#v", action))
		}
	}
}

func printCommitMessage(cmd *cobra.Command, completion string) {
	cmd.Printf("💬 Commit message:"+
		"\n==================================================================================================\n%s"+
		"\n==================================================================================================\n",
		completion,
	)
}

// clearScreen clears the terminal screen. It uses platform-specific
// commands based on the runtime environment.
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}
}
