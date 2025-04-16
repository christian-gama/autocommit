package openai

// askConfigsCliImpl is the implementation of AskConfigsCli.
// type askConfigsCliImpl struct{}
//
// // Execute asks the user for the configuration.
// func (a *askConfigsCliImpl) Execute() (llm.Config, error) {
// 	provider := NewOpenAIProvider()
//
// 	questions := llm.CreateQuestions(
// 		provider,
// 		a.createApiKeyQuestion,
// 		a.createModelQuestion,
// 	)
//
// 	type Answers struct {
// 		Model        string
// 		OpenAIAPIKey string
// 	}
//
// 	var answers Answers
//
// 	err := survey.Ask(questions, &answers)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return NewConfig(answers.OpenAIAPIKey, answers.Model), nil
// }

// func (a *askConfigsCliImpl) createModelQuestion(provider llm.Provider) *survey.Question {
// 	prompt := survey.Select{
// 		Message: "Model name",
// 		Help:    "A model can be an algorithm or a set of algorithms that have been trained on data to make predictions or decisions.",
// 		Default: GPT4oMini,
// 		Options: AllowedModels,
// 		VimMode: true,
// 	}
//
// 	return &survey.Question{
// 		Name:   "Model",
// 		Prompt: &prompt,
// 	}
// }
//
// // func (a *askConfigsCliImpl) createApiKeyQuestion(provider llm.Provider) *survey.Question {
// // 	prompt := survey.Password{
// // 		Message: "OpenAI API Key",
// // 		Help:    "The OpenAI API Key is used to authenticate your requests to the OpenAI API.",
// // 	}
// //
// // 	return &survey.Question{
// // 		Name:   "OpenAIAPIKey",
// // 		Prompt: &prompt,
// // 		Validate: func(ans interface{}) error {
// // 			apiKey, ok := ans.(string)
// // 			if !ok {
// // 				return errors.New("invalid API key")
// // 			}
// // 			return ValidateApiKey(apiKey, provider)
// // 		},
// // 	}
// // }
//
// // NewAskConfigsCli creates a new instance of AskConfigsCli.
// func NewAskConfigsCli() llm.AskConfigsCli {
// 	return &askConfigsCliImpl{}
// }
//
// type AskToChangeModelCli interface {
// 	Execute() (bool, error)
// }
//
// type askToChangeModelCliImpl struct{}
//
// func (a *askToChangeModelCliImpl) Execute() (bool, error) {
// 	var provider llm.Provider
// 	questions := llm.CreateQuestions(
// 		provider,
// 		a.createModelQuestion,
// 	)
//
// 	type Answers struct {
// 		ChangeModel bool
// 	}
//
// 	var answers Answers
//
// 	err := survey.Ask(questions, &answers)
// 	if err != nil {
// 		return false, err
// 	}
//
// 	return answers.ChangeModel, nil
// }
//
// func (a *askToChangeModelCliImpl) createModelQuestion(provider llm.Provider) *survey.Question {
// 	prompt := survey.Confirm{
// 		Message: "You reached the maximum number of tokens, but there is a model that can generate longer messages. Do you want to temporarily change the model?",
// 		Help:    "A model have a limited amount of tokens that can be generated at once. If you want to generate longer messages, you can temporarily change the model.",
// 		Default: true,
// 	}
//
// 	return &survey.Question{
// 		Name:   "ChangeModel",
// 		Prompt: &prompt,
// 	}
// }
//
// func NewAskToChangeModelCli() AskToChangeModelCli {
// 	return &askToChangeModelCliImpl{}
// }
