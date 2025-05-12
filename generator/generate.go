package generator

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/christian-gama/autocommit/git"
	"github.com/christian-gama/autocommit/instruction"
	"github.com/tmc/langchaingo/llms"
)

type Generator struct {
	model llms.Model
	msgs  []llms.MessageContent
}

func New(model llms.Model) (*Generator, error) {
	instruction, err := instruction.Load()
	if err != nil {
		return nil, err
	}

	content, err := getContent()
	if err != nil {
		return nil, err
	}

	msgs := make([]llms.MessageContent, 0, len(instruction)+1)

	msgs = append(msgs, llms.MessageContent{
		Role:  llms.ChatMessageTypeSystem,
		Parts: []llms.ContentPart{llms.TextContent{Text: instruction}},
	})

	msgs = append(msgs, llms.MessageContent{
		Role:  llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{llms.TextContent{Text: content}},
	})

	return &Generator{
		model: model,
		msgs:  msgs,
	}, nil
}

func (a *Generator) Generate(ctx context.Context, additionalPrompts ...string) (string, error) {
	if len(additionalPrompts) > 0 {
		for _, prompt := range additionalPrompts {
			msg := llms.MessageContent{
				Role:  llms.ChatMessageTypeHuman,
				Parts: []llms.ContentPart{llms.TextContent{Text: prompt}},
			}
			a.msgs = append(a.msgs, msg)
		}
	}

	resp, err := a.model.GenerateContent(ctx, a.msgs)
	if err != nil {
		return "", err
	}

	choices := resp.Choices
	if len(choices) < 1 {
		return "", errors.New("empty response from model")
	}

	c1 := choices[0]

	a.msgs = append(a.msgs, llms.MessageContent{
		Role:  llms.ChatMessageTypeAI,
		Parts: []llms.ContentPart{llms.TextContent{Text: c1.Content}},
	})

	return strings.TrimSpace(c1.Content), nil
}

func getContent() (string, error) {
	diff, err := git.MinimalDiff()
	if err != nil {
		return "", err
	}

	files, err := git.ListFiles()
	if err != nil {
		return "", err
	}

	prompt := fmt.Sprintf(
		"Here is the current project structure:\n%s\n\nAnd finally the git diff output:\n%s\n",
		files,
		diff,
	)

	return prompt, nil
}
