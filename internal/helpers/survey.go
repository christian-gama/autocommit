package helpers

import "github.com/AlecAivazis/survey/v2"

type SurveyFunc func() *survey.Question

func CreateQuestions(funcs ...SurveyFunc) []*survey.Question {
	var questions []*survey.Question

	for _, f := range funcs {
		questions = append(questions, f())
	}

	return questions
}
