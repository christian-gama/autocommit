package openai

import (
	"context"
	"errors"
	"time"

	"github.com/sashabaranov/go-openai"
)

// Chat is the interface that wraps the Response method.
type Chat interface {
	// Response returns the response from the AI.
	Response(config *Config, system *System, input string) (string, error)
}

// chatImpl is an implementation of Chat.
type chatImpl struct {
	repo ConfigRepo

	messages []openai.ChatCompletionMessage
}

// Chat implements the Chat interface.
func (c *chatImpl) Response(config *Config, system *System, input string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 18*time.Second)
	defer cancel()

	response, err := openai.
		NewClient(config.ApiKey).
		CreateChatCompletion(
			ctx,
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

	c.messages = append(c.messages, response.Choices[0].Message)

	return response.Choices[0].Message.Content, nil
}

func (c *chatImpl) createMessages(system *System, userInput string) []openai.ChatCompletionMessage {
	if len(c.messages) == 0 && system != nil {
		c.messages = append(c.messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: system.Message,
			Name:    system.Name,
		})
	}

	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: userInput,
		Name:    "UserInput",
	})

	return c.messages
}

func (c *chatImpl) checkError(err error) error {
	if err == nil {
		return nil
	}

	return err
}

// NewChat creates a new instance of Chat.
func NewChat(repo ConfigRepo) Chat {
	return &chatImpl{
		repo:     repo,
		messages: make([]openai.ChatCompletionMessage, 0),
	}
}
