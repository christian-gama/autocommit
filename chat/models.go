package chat

import openai "github.com/sashabaranov/go-openai"

var Models = []string{
	openai.GPT3Ada,
	openai.GPT3Babbage,
	openai.GPT3Curie,
	openai.GPT3CurieInstructBeta,
	openai.GPT3Davinci,
	openai.GPT3DavinciInstructBeta,
	openai.GPT3Dot5Turbo,
	openai.GPT3Dot5Turbo0301,
	openai.GPT3TextAda001,
	openai.GPT3TextBabbage001,
	openai.GPT3TextCurie001,
	openai.GPT3TextDavinci001,
	openai.GPT3TextDavinci002,
	openai.GPT3TextDavinci003,
	openai.GPT4,
	openai.GPT40314,
	openai.GPT432K,
	openai.GPT432K0314,
}
