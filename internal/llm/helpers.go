package llm

import (
	"github.com/AlecAivazis/survey/v2"
)

type surveyFunc func(Provider) *survey.Question

// CreateQuestions creates a list of questions.
func CreateQuestions(provider Provider, funcs ...surveyFunc) []*survey.Question {
	var questions []*survey.Question

	for _, f := range funcs {
		questions = append(questions, f(provider))
	}

	return questions
}
