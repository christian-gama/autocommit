package openai

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
)

type Chat interface {
	Response(system *System, input string) (string, error)
}

type chatImpl struct {
	repo Repo
}

func (c *chatImpl) Response(system *System, input string) (string, error) {
	config, err := c.repo.GetConfig()
	if err != nil {
		return "", err
	}

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

func NewChat(repo Repo) Chat {
	return &chatImpl{
		repo: repo,
	}
}
