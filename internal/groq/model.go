package groq

const (
	DEEPSEEK        = "deepseek-r1-distill-llama-70b"
	LLAMA           = "llama3-70b-8192"
	LLAMA_VERSATILE = "llama-3.3-70b-versatile"
	GEMMA           = "gemma2-9b-it"
)

var AllowedModels = []string{
	DEEPSEEK,
	LLAMA,
	LLAMA_VERSATILE,
	GEMMA,
}
