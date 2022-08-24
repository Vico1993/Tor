package handlers

import (
	"strings"

	"github.com/Vico1993/Tor/domain/hue"
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

	// Create default message
	msg := tgbotapi.NewMessage(
		update.CallbackQuery.Message.Chat.ID,
		"I didn't understand...",
	)

	// Manage the callbacy query
	switch tmp[0] {
		case "off":
			msg.Text = "Shutting down: " + tmp[1]
			hue.ShutDownGroup(tmp[1])
		case "on":
			msg.Text = "Powering up: " + tmp[1]
			hue.PowerUpGroup(tmp[1])
	}

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}