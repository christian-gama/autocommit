package provider

import (
	"context"
	"testing"

	"github.com/tmc/langchaingo/llms"
)

func TestFlattenTextMessages(t *testing.T) {
	message := flattenTextMessages([]llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{llms.TextPart("Use conventional commits.")},
		},
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.TextPart("Diff goes here.")},
		},
		{
			Role:  llms.ChatMessageTypeAI,
			Parts: []llms.ContentPart{llms.TextPart("feat: add thing")},
		},
	})

	if message.Role != llms.ChatMessageTypeHuman {
		t.Fatalf("role = %q, want %q", message.Role, llms.ChatMessageTypeHuman)
	}

	text, ok := message.Parts[0].(llms.TextContent)
	if !ok {
		t.Fatalf("part = %T, want llms.TextContent", message.Parts[0])
	}

	const want = "System:\nUse conventional commits.\n\nUser:\nDiff goes here.\n\nAssistant:\nfeat: add thing"
	if text.Text != want {
		t.Fatalf("text = %q, want %q", text.Text, want)
	}
}

func TestGoogleAIModelFlattensMultipleMessages(t *testing.T) {
	recorder := &recordingModel{}
	model := &googleAIModel{model: recorder}

	_, err := model.GenerateContent(context.Background(), []llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{llms.TextPart("System prompt")},
		},
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.TextPart("User prompt")},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(recorder.messages) != 1 {
		t.Fatalf("message count = %d, want 1", len(recorder.messages))
	}

	text, ok := recorder.messages[0].Parts[0].(llms.TextContent)
	if !ok {
		t.Fatalf("part = %T, want llms.TextContent", recorder.messages[0].Parts[0])
	}

	const want = "System:\nSystem prompt\n\nUser:\nUser prompt"
	if text.Text != want {
		t.Fatalf("text = %q, want %q", text.Text, want)
	}
}

func TestGoogleAIModelKeepsSingleMessage(t *testing.T) {
	recorder := &recordingModel{}
	model := &googleAIModel{model: recorder}

	_, err := model.GenerateContent(context.Background(), []llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.TextPart("User prompt")},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	text := recorder.messages[0].Parts[0].(llms.TextContent)
	if text.Text != "User prompt" {
		t.Fatalf("text = %q, want %q", text.Text, "User prompt")
	}
}

type recordingModel struct {
	messages []llms.MessageContent
}

func (m *recordingModel) GenerateContent(
	_ context.Context,
	messages []llms.MessageContent,
	_ ...llms.CallOption,
) (*llms.ContentResponse, error) {
	m.messages = messages
	return &llms.ContentResponse{
		Choices: []*llms.ContentChoice{{Content: "commit message"}},
	}, nil
}

func (m *recordingModel) Call(_ context.Context, _ string, _ ...llms.CallOption) (string, error) {
	return "commit message", nil
}
