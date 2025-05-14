package cli

import (
	"github.com/christian-gama/autocommit/v2/ask"
	"github.com/christian-gama/autocommit/v2/config"
	"github.com/christian-gama/autocommit/v2/generator"
	"github.com/christian-gama/autocommit/v2/instruction"
)

var (
	_config    *config.Config
	_generator *generator.Generator
	_askAction *ask.Action
	_askConfig *ask.Config
)

var instructionFlag string

func init() {
	AutoCommitCmd.Flags().
		StringVarP(&instructionFlag, "instruction", "i", "", "Add custom instruction to guide commit message")

	if !instruction.HasInstruction() {
		if err := instruction.Create(); err != nil {
			panic(err)
		}
	}

	cfg, err := config.LoadOrNew()
	if err != nil {
		panic(err)
	}

	_config = cfg
	_askAction = ask.NewAction()
	_askConfig = ask.NewConfig()

	AutoCommitCmd.AddCommand(configureCmd)
	AutoCommitCmd.AddCommand(instructionsCmd)
	instructionsCmd.AddCommand(restoreInstructionsCmd)
}
