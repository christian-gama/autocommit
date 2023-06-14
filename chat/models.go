package chat

const (
	GPT3Dot5Turbo    = "gpt3.5-turbo"
	GPT3Dot5Turbo16k = "gpt-3.5-turbo-16k"
	GPT4             = "gpt-4"
	GPT432K          = "gpt-4-32k"
)

// Model is a map of model names to their instructions.
type Model map[string]string

var ModelMap = Model{
	"gpt-3.5-turbo":     MinimalInstructions,
	"gpt-3.5-turbo-16k": DetailedInstructions,
	"gpt-4":             MinimalInstructions,
	"gpt-4-32k":         DetailedInstructions,
}
