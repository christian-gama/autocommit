package openai

//
// import (
// 	"github.com/christian-gama/autocommit/internal/llm"
// 	"github.com/christian-gama/autocommit/internal/storage"
// )
//
// func MakeConfigRepo() llm.ConfigRepo {
// 	return NewConfigRepo(storage.NewStorage("config.json"))
// }
//
// func MakeAskConfigsCli() llm.AskConfigsCli {
// 	return NewAskConfigsCli()
// }
//
// func MakeVerifyConfigCommand() llm.VerifyConfigCommand {
// 	return NewVerifyConfigCommand(MakeConfigRepo())
// }
//
// func MakeResetConfigCommand() ResetConfigCommand {
// 	return NewResetConfigCommand(MakeConfigRepo())
// }
//
// func MakeUpdateConfigCommand() llm.UpdateConfigCommand {
// 	return NewUpdateConfigCommand(MakeConfigRepo())
// }
//
// func MakeAskToChangeModelCli() AskToChangeModelCli {
// 	return NewAskToChangeModelCli()
// }
//
// func MakeChatCommand() llm.ChatCommand {
// 	return NewChatCommand(NewChat(MakeConfigRepo()))
// }
