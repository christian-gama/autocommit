package provider

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/christian-gama/autocommit/v2/config"
	"github.com/tmc/langchaingo/llms"
	lcgoogleai "github.com/tmc/langchaingo/llms/googleai"
)

type GoogleAI struct{}

func (g GoogleAI) New(cfg *config.Config) (llms.Model, error) {
	llm, ok := cfg.LLM(g.Name())
	if !ok {
		return nil, fmt.Errorf("no Google AI LLM provider found")
	}

	if err := os.Setenv("API_KEY", llm.Credential); err != nil {
		return nil, fmt.Errorf("set API_KEY: %w", err)
	}

	model, err := lcgoogleai.New(
		context.Background(),
		lcgoogleai.WithAPIKey(llm.Credential),
		lcgoogleai.WithDefaultModel(llm.Model),
	)
	if err != nil {
		return nil, err
	}

	return &googleAIModel{model: model}, nil
}

type googleAIModel struct {
	model llms.Model
}

func (m *googleAIModel) GenerateContent(
	ctx context.Context,
	messages []llms.MessageContent,
	options ...llms.CallOption,
) (*llms.ContentResponse, error) {
	if len(messages) > 1 {
		messages = []llms.MessageContent{flattenTextMessages(messages)}
	}

	return m.model.GenerateContent(ctx, messages, options...)
}

func (m *googleAIModel) Call(
	ctx context.Context,
	prompt string,
	options ...llms.CallOption,
) (string, error) {
	return m.model.Call(ctx, prompt, options...)
}

func flattenTextMessages(messages []llms.MessageContent) llms.MessageContent {
	var prompt strings.Builder

	for i, message := range messages {
		if i > 0 {
			prompt.WriteString("\n\n")
		}

		prompt.WriteString(roleLabel(message.Role))
		prompt.WriteString(":\n")

		for j, part := range message.Parts {
			if j > 0 {
				prompt.WriteString("\n")
			}
			fmt.Fprint(&prompt, part)
		}
	}

	return llms.MessageContent{
		Role:  llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{llms.TextPart(prompt.String())},
	}
}

func roleLabel(role llms.ChatMessageType) string {
	switch role {
	case llms.ChatMessageTypeSystem:
		return "System"
	case llms.ChatMessageTypeAI:
		return "Assistant"
	case llms.ChatMessageTypeHuman:
		return "User"
	default:
		return "Message"
	}
}

func (g GoogleAI) Name() string {
	return "Google AI"
}

func (g GoogleAI) Models() []string {
	return []string{
		"gemini-flash-latest",
		"gemini-flash-lite-latest",
		"gemini-3-pro-preview",
	}
}
