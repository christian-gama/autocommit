package cli

import (
	"github.com/christian-gama/autocommit/v2/ask"
	"github.com/christian-gama/autocommit/v2/config"
	"github.com/christian-gama/autocommit/v2/generator"
	"github.com/christian-gama/autocommit/v2/instruction"
)

var _config *config.Config
var _generator *generator.Generator
var _askAction *ask.Action
var _askConfig *ask.Config

func init() {
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
