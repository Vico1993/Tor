package handlers

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleReply(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	callback := tgbotapi.NewCallback(
		update.CallbackQuery.ID,
		update.CallbackQuery.Data,
	)

	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}

	tmp := strings.Split(update.CallbackQuery.Data, " - ")

	botRespond(botParameter{
		ChatId: update.CallbackQuery.Message.Chat.ID,
		Bot: bot,
		Command: tmp[0],
		CommandParam: tmp[1],
	})
}