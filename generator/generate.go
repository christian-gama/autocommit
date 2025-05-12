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

	return strings.TrimSpace(boundText(c1.Content)), nil
}

const maxLength = 100

func boundText(text string) string {
	// Split while preserving empty lines
	lines := strings.Split(text, "\n")

	var result []string
	var lastLineWasEmpty = false

	for _, line := range lines {
		// Handle empty lines (newlines)
		if len(strings.TrimSpace(line)) == 0 {
			// If this is an empty line but the previous wasn't empty
			// we add exactly one empty line to preserve a single newline
			if !lastLineWasEmpty {
				result = append(result, "")
				lastLineWasEmpty = true
			}
			// Skip additional empty lines to avoid multiple consecutive newlines
			continue
		}

		// Reset the empty line tracker
		lastLineWasEmpty = false

		// Process non-empty line until it's fully bounded
		currentLine := line
		for len(currentLine) > 0 {
			if len(currentLine) <= maxLength {
				// Line fits within the limit
				result = append(result, currentLine)
				break
			} else {
				// Need to truncate and wrap
				// Find the last space before the limit if possible
				cutPoint := maxLength - 1 // Leave room for the hyphen

				// Try to find a space to break at
				foundSpace := false
				for i := cutPoint; i >= 0; i-- {
					if currentLine[i] == ' ' {
						cutPoint = i
						foundSpace = true
						break
					}
				}

				// If no space was found, just cut at the limit
				if !foundSpace {
					result = append(result, currentLine[:cutPoint]+"-")
					currentLine = currentLine[cutPoint:]
				} else {
					// Cut at the space
					result = append(result, currentLine[:cutPoint])
					currentLine = currentLine[cutPoint+1:] // Skip the space
				}
			}
		}
	}

	return strings.Join(result, "\n")
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
