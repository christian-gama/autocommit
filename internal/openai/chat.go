package openai

import (
	"context"
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// Chat is the interface that wraps the Response method.
type Chat interface {
	// Response returns the response from the AI.
	Response(config *Config, system *System, input string) (string, error)
}

// chatImpl is an implementation of Chat.
type chatImpl struct {
	repo Repo
}

// Chat implements the Chat interface.
func (c *chatImpl) Response(config *Config, system *System, input string) (string, error) {
	fmt.Printf("🤖 Using model: %s\n", config.Model)
	fmt.Printf("⌛ Waiting for response from OpenAI...\n\n")

	response, err := openai.
		NewClient(config.ApiKey).
		CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:       config.Model,
				Temperature: config.Temperature,
				Messages:    c.createMessages(system, input),
			},
		)

	if err != nil {
		return "", c.checkError(err)
	}

	if len(response.Choices) == 0 {
		return "", errors.New("Received empty response from AI")
	}

	return response.Choices[0].Message.Content, nil
}

func (c *chatImpl) createMessages(system *System, userInput string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: system.Message,
			Name:    system.Name,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
			Name:    "UserInput",
		},
	}
}

func (c *chatImpl) checkError(err error) error {
	if err == nil {
		return nil
	}

	return err
}

// NewChat creates a new instance of Chat.
func NewChat(repo Repo) Chat {
	return &chatImpl{
		repo: repo,
	}
}
