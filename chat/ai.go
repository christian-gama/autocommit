package chat

import (
	"context"
	"errors"

	openai "github.com/sashabaranov/go-openai"
)

// Config is the configuration for the chat service.
type ChatService struct {
	config *Config
	client *openai.Client
}

// NewChatService creates a new chat service.
func NewChatService(config *Config) *ChatService {
	return &ChatService{
		config: config,
		client: openai.NewClient(config.APIKey),
	}
}

// GenerateChatResponse generates a chat response based on the user input.
func (cs *ChatService) GenerateChatResponse(userInput string) (string, error) {
	response, err := cs.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       cs.config.Model,
			Temperature: cs.config.Temperature,
			Messages:    cs.createMessages(userInput),
		},
	)

	if err != nil {
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", errors.New("Received empty response from AI")
	}

	return response.Choices[0].Message.Content, nil
}

func (cs *ChatService) createMessages(userInput string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: cs.config.Template,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		},
	}
}
