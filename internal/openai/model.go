package openai

import "github.com/sashabaranov/go-openai"

const (
	O1Mini       = openai.O1Mini
	O3Mini       = openai.O3Mini
	O4Mini       = openai.O4Mini
	GPT4o        = openai.GPT4o
	GPT4oLatest  = openai.GPT4oLatest
	GPT4oMini    = openai.GPT4oMini
	GPT4Turbo    = openai.GPT4Turbo
	GPT4Dot1     = openai.GPT4Dot1
	GPT4Dot1Mini = openai.GPT4Dot1Mini
)

var AllowedModels = []string{
	O1Mini,
	O3Mini,
	O4Mini,
	GPT4o,
	GPT4oLatest,
	GPT4oMini,
	GPT4Turbo,
	GPT4Dot1,
	GPT4Dot1Mini,
}
