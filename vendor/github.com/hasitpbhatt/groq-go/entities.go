package groq

import "net/http"

// Client represents a client for interacting with the Groq API.
type Client struct {
	// apiKey is the API key for authentication.
	apiKey string
	// chatCompletionURL is the endpoint for chat completions.
	chatCompletionURL string
	// httpClient is the HTTP client used for making requests.
	httpClient *http.Client
}

// Message represents a single message in the chat completion request.
// It contains the role of the message sender (e.g., user or system) and the content of the message itself.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Option represents a function that modifies the requestBody.
type Option func(*requestBody)

// ClientOption represents a function that modifies the Client.
type ClientOption func(*Client)

// WithAPIKey sets the API key for the client.
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

type requestBody struct {
	// Messages represents a slice of Message structures for the chat completion request.
	Messages []Message `json:"messages"`
	// Model specifies the model to use for the chat completion.
	Model string `json:"model"`
	// MaxTokens sets the maximum number of tokens to generate.
	MaxTokens int `json:"max_tokens"`
	// ResponseFormat specifies the format of the response.
	ResponseFormat struct {
		Type string `json:"type,omitempty"`
	} `json:"response_format,omitempty"`
	// Seed sets the seed for the random number generator.
	Seed int `json:"seed,omitempty"`
	// Stream indicates whether to stream the response.
	Stream bool `json:"stream"`
	// Stop specifies the sequence where the text generation should stop.
	Stop *string `json:"stop,omitempty"`
	// Temperature controls randomness in the output.
	Temperature float64 `json:"temperature"`
	// TopP controls the diversity of the output.
	TopP float64 `json:"top_p"`
}

// ChatCompletionResponse represents the structure of the response received from the Groq API for chat completions.
// It contains the ID of the completion, the object type, the creation time, the model used, the choices made, the usage statistics, the system fingerprint, and the x_groq information.
type ChatCompletionResponse struct {
	// ID represents the unique identifier for the chat completion response.
	ID string `json:"id,omitempty"`
	// Object specifies the type of object returned in the response.
	Object string `json:"object,omitempty"`
	// Created indicates the timestamp when the response was created.
	Created int `json:"created,omitempty"`
	// Model specifies the model used for the chat completion.
	Model string `json:"model,omitempty"`
	// Choices represents a slice of choice structures containing information about each choice.
	Choices []struct {
		// Index specifies the index of the choice.
		Index int `json:"index,omitempty"`
		// Message contains the message content of the choice.
		Message Message `json:"message,omitempty"`
		// Logprobs represents the log probabilities of the choice.
		Logprobs interface{} `json:"logprobs,omitempty"`
		// FinishReason indicates the reason why the choice was finished.
		FinishReason string `json:"finish_reason,omitempty"`
	} `json:"choices,omitempty"`
	// Usage contains usage statistics for the chat completion.
	Usage struct {
		// QueueTime specifies the time spent in the queue.
		QueueTime float64 `json:"queue_time,omitempty"`
		// PromptTokens indicates the number of tokens in the prompt.
		PromptTokens int `json:"prompt_tokens,omitempty"`
		// PromptTime specifies the time spent processing the prompt.
		PromptTime float64 `json:"prompt_time,omitempty"`
		// CompletionTokens indicates the number of tokens in the completion.
		CompletionTokens int `json:"completion_tokens,omitempty"`
		// CompletionTime specifies the time spent generating the completion.
		CompletionTime float64 `json:"completion_time,omitempty"`
		// TotalTokens indicates the total number of tokens processed.
		TotalTokens int `json:"total_tokens,omitempty"`
		// TotalTime specifies the total time spent processing the request.
		TotalTime float64 `json:"total_time,omitempty"`
	} `json:"usage,omitempty"`
	// SystemFingerprint represents a unique identifier for the system.
	SystemFingerprint string `json:"system_fingerprint,omitempty"`
	// XGroq contains additional information about the Groq system.
	XGroq struct {
		// ID specifies the unique identifier for the Groq system.
		ID string `json:"id,omitempty"`
	} `json:"x_groq,omitempty"`
}
