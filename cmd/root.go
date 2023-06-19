package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/christian-gama/autocommit/internal/autocommit"
	"github.com/christian-gama/autocommit/internal/openai"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "autocommit",
	Run:   runCmd,
	Short: "Autocommit is a CLI tool that uses OpenAI's models to generate commit messages based on the changes made in the repository.",
	ValidArgs: []string{
		"set",
		"reset",
	},
}

// Execute executes the root command.
func Execute() error {
	return cmd.Execute()
}

func runCmd(cmd *cobra.Command, args []string) {
	var err error
	config, err = verifyConfigCommand.Execute(askConfigsCli.Execute)
	if err != nil {
		panic(err)
	}

	fmt.Printf("🤖 Using model: %s\n", config.Model)

	handleCmd(cmd, args)
}

func handleCmd(cmd *cobra.Command, args []string) {
	response, err := generatorCommand.Execute(config)
	if err != nil {
		handleMaxToken(err, cmd, args)
	}

	fmt.Printf("📝 Commit message generated: \n%s\n\n", response)

	handlePostCommit(response, cmd, args)
}

func handlePostCommit(response string, cmd *cobra.Command, args []string) {
	option, err := postCommitCli.Execute()
	if err != nil {
		panic(err)
	}

	switch option {
	case autocommit.CommitChangesOption:
		if err := commitCommand.Execute(response); err != nil {
			panic(err)
		}

	case autocommit.CopyToClipboardOption:
		if err := clipboardCommand.Execute(response); err != nil {
			panic(err)
		}

	case autocommit.RegenerateOption:
		handleCmd(cmd, args)

	case autocommit.AddInstructionOption:
		instructions, err := addInstructionCli.Execute()
		if err != nil {
			panic(err)
		}

		if err := addInstructionCommand.Execute(config, nil, instructions); err != nil {
			panic(err)
		}

		handleCmd(cmd, args)

	case autocommit.ExitOption:
		os.Exit(0)
	}
}

func handleMaxToken(err error, cmd *cobra.Command, args []string) {
	var isTokenError = strings.Contains(err.Error(), "Please reduce the length of the messages")
	var isElegibleModel = config.Model == openai.GPT3Dot5Turbo || config.Model == openai.GPT4

	if !isTokenError || !isElegibleModel {
		panic(err)
	}

	var modelMap = map[string]string{
		openai.GPT3Dot5Turbo: openai.GPT3Dot5Turbo16k,
		openai.GPT4:          openai.GPT432K,
	}

	answer, err := askToChangeModelCli.Execute()
	if err != nil {
		panic(err)
	}

	if answer {
		config.Model = modelMap[config.Model]
		handleCmd(cmd, args)
		return
	}

	panic(
		"🚧 You reached the maximum allowed token for this model. You can try a new model running 'autocommit set --model <model>' or decrease the amount of files being commited.",
	)
}
