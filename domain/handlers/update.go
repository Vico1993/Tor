package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handlerContext struct {
	ChatId 			int64
	Command 		string
	CommandParam 	string

}

func buildUpdateContext(update tgbotapi.Update, bot *tgbotapi.BotAPI) handlerContext {
	return handlerContext{
		ChatId: update.Message.Chat.ID,
		Command: update.Message.Command(),
		CommandParam: "",
	}
}
