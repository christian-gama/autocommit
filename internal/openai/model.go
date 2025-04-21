package openai

const (
	O1Mini      = "o1-mini"
	GPT4o       = "gpt-4o"
	GPT4oLatest = "chatgpt-4o-latest"
	GPT4oMini   = "gpt-4o-mini"
	GPT4Turbo   = "gpt-4-turbo"
)

var AllowedModels = []string{
	O1Mini,
	GPT4o,
	GPT4oLatest,
	GPT4oMini,
	GPT4Turbo,
}
