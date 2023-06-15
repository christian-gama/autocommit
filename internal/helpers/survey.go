package helpers

import "github.com/AlecAivazis/survey/v2"

type surveyFunc func() *survey.Question

// CreateQuestions creates a list of questions.
func CreateQuestions(funcs ...surveyFunc) []*survey.Question {
	var questions []*survey.Question

	for _, f := range funcs {
		questions = append(questions, f())
	}

	return questions
}
