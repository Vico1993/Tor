package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// We are checking if the update we receive is a Message
// From a Group
// if it's a Command
func shouldReact(update tgbotapi.Update) bool {
	return update.Message == nil &&
		update.FromChat().Type != "group" &&
		!update.Message.IsCommand()
}

func HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if !shouldReact(update) {
		// TODO: Log this why we don't act upon it
		return
	}

	// Create default message
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I didn't understand...")

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}