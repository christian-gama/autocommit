package provider

import (
	"github.com/christian-gama/autocommit/config"
	"github.com/tmc/langchaingo/llms"
)

type Func func(*config.Config) (llms.Model, error)
