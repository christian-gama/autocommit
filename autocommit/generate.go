package autocommit

import (
	"context"
	"errors"
	"fmt"

	"github.com/christian-gama/autocommit/git"
	"github.com/christian-gama/autocommit/systemmsg"
	"github.com/tmc/langchaingo/llms"
)

type Autocommit struct {
	model llms.Model
	msgs  []llms.MessageContent
}

func New(model llms.Model) (*Autocommit, error) {

	systemMsg, err := systemmsg.Load()
	if err != nil {
		return nil, err
	}

	additionalContextMsg, err := addContextMsg()
	if err != nil {
		return nil, err
	}

	msgs := make([]llms.MessageContent, 0, len(systemMsg)+1)

	msgs = append(msgs, llms.MessageContent{
		Role:  llms.ChatMessageTypeSystem,
		Parts: []llms.ContentPart{llms.TextContent{Text: systemMsg}},
	})

	msgs = append(msgs, llms.MessageContent{
		Role:  llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{llms.TextContent{Text: additionalContextMsg}},
	})

	return &Autocommit{
		model: model,
		msgs:  msgs,
	}, nil
}

func (a *Autocommit) Generate(ctx context.Context, additionalPrompts ...string) (string, error) {
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

	return c1.Content, nil
}

func addContextMsg() (string, error) {
	diff, err := git.MinimalDiff()
	if err != nil {
		return "", err
	}

	files, err := git.ListFiles()
	if err != nil {
		return "", err
	}

	lastMessages, err := git.LastCommitMessages(8)
	if err != nil {
		return "", err
	}

	prompt := fmt.Sprintf(
		"Sample of previous git messages so that you can keep new messages style consistent, such as casing, spacing and line breaks:\n%s\nNow here is the current project structure:\n%s\nAnd finally the git diff output:\n%s",
		lastMessages,
		files,
		diff,
	)

	return prompt, nil
}
