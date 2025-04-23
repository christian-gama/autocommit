package groq

import (
	"errors"

	"github.com/christian-gama/autocommit/internal/llm"
	groq "github.com/hasitpbhatt/groq-go"
)

// chatImpl is an implementation of Chat.
// chatImpl is an implementation of Chat.
type chatImpl struct {
	repo     llm.ConfigRepo
	messages []groq.Message
}

// Chat implements the Chat interface.
func (c *chatImpl) Response(config llm.Config, system *llm.System, input string) (string, error) {
	client := groq.NewClient(groq.WithAPIKey(config.GetAPIKey()))

	c.messages = c.createMessages(system, input)
	response, err := client.ChatCompletion(c.messages, groq.WithModel(config.GetModel()))
	if err != nil {
		return "", c.checkError(err)
	}

	if len(response.Choices) == 0 {
		return "", errors.New("received empty response from Groq LLM")
	}

	c.messages = append(c.messages, response.Choices[0].Message)

	return response.Choices[0].Message.Content, nil
}

func (c *chatImpl) createMessages(system *llm.System, userInput string) []groq.Message {
	if len(c.messages) == 0 && system != nil {
		c.messages = append(c.messages, groq.Message{
			Role:    "system",
			Content: system.Message,
		})
	}

	c.messages = append(c.messages, groq.Message{
		Role:    "user",
		Content: userInput,
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
func NewGroqChat(repo llm.ConfigRepo) llm.Chat {
	return &chatImpl{
		repo:     repo,
		messages: make([]groq.Message, 0),
	}
}
