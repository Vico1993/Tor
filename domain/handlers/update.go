package handlers

import (
	"github.com/Vico1993/Tor/domain/hue"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// We are checking if the update we receive is a Message
// From a Group
// if it's a Command
func shouldNotAct(update tgbotapi.Update) bool {
	return update.FromChat().Type != "group" || update.Message == nil || !update.Message.IsCommand()
}

func HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if shouldNotAct(update) {
		// TODO: Log this why we don't act upon it
		return
	}

	// Create default message
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I didn't understand...")

	switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "off":
			msg.Text = "let's go to bed!"
			hue.SetLight(false)
		case "on":
			msg.Text = "Light is coming"
			hue.SetLight(true)
		default:
			msg.Text = "I don't know that command"
	}

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}