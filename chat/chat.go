package chat

import (
	"context"
	"errors"
	"fmt"
	"strings"

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

func (cs *ChatService) createMessages(userInput string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: cs.config.Template,
			Name:    "GitCommitMessageGenerator",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("Output of 'git diff --cached':\n\"%s\"", userInput),
			Name:    "UserInput",
		},
	}
}

// Response generates a chat response based on the user input.
func (cs *ChatService) Response(userInput string) (string, error) {
	response, err := cs.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       cs.config.ModelName,
			Temperature: cs.config.Temperature,
			Messages:    cs.createMessages(userInput),
		},
	)

	if err != nil {
		return "", checkError(err)
	}

	if len(response.Choices) == 0 {
		return "", errors.New("Received empty response from AI")
	}

	return response.Choices[0].Message.Content, nil
}

func checkError(err error) error {
	if strings.Contains(err.Error(), "does not exist") {
		return errors.New("The model you specified does not exist or is not available for your account.")
	}

	return err
}
