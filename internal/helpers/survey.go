// package helpers
//
// import (
// 	"github.com/AlecAivazis/survey/v2"
// 	"github.com/christian-gama/autocommit/internal/llm"
// )
//
// type surveyFunc func(llm.Provider) *survey.Question
//
// // CreateQuestions creates a list of questions.
// func CreateQuestions(provider llm.Provider, funcs ...surveyFunc) []*survey.Question {
// 	var questions []*survey.Question
//
// 	for _, f := range funcs {
// 		questions = append(questions, f(provider))
// 	}
//
// 	return questions
// }
