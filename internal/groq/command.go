package groq

import (
	"errors"

	"github.com/christian-gama/autocommit/internal/llm"
)

type chatCommandImpl struct {
	chat llm.Chat
}

// Execute implements the ChatCommand interface.
func (c *chatCommandImpl) Execute(
	config llm.Config,
	system *llm.System,
	input string,
) (string, error) {
	openaiConfig, ok := config.(*GroqConfig)
	if !ok {
		return "", errors.New("invallid config type")
	}
	response, err := c.chat.Response(openaiConfig, system, input)
	if err != nil {
		return "", err
	}
	return response, nil
}

// NewChatCommand creates a new instance of ChatCommand.
func NewChatCommand(chat llm.Chat) llm.ChatCommand {
	return &chatCommandImpl{
		chat: chat,
	}
}
