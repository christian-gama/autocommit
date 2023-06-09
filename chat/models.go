package chat

import openai "github.com/sashabaranov/go-openai"

var Models = []string{
	openai.GPT3Dot5Turbo,
	openai.GPT4,
	openai.GPT432K,
}
