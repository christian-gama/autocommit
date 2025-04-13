package llm

// Chat is the interface that wraps the Response method.
type Chat interface {
	// Response returns the response from the AI.
	Response(config Config, system *System, input string) (string, error)
}
