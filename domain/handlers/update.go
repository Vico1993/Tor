package handlers

import (
	"github.com/Vico1993/Tor/domain/hue"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot should respond
func shouldNotAct(update tgbotapi.Update) bool {
	// If it's not a Group chat
	if update.FromChat().Type != "group" {
		return false
	}

	// If it's not a Message or CallBackQuery
	if update.Message == nil && update.CallbackQuery == nil {
		return false
	}

	// If it's a message and not an Command
	return update.Message != nil && !update.Message.IsCommand()
}

func listGroup(command string) [][]tgbotapi.InlineKeyboardButton {
	groups, err := hue.GetAllGroup()
	if err != nil {
		panic(err)
	}

	var buttons [][]tgbotapi.InlineKeyboardButton
	var line []tgbotapi.InlineKeyboardButton
	for key, group := range groups {
		line = append(
			line,
			tgbotapi.NewInlineKeyboardButtonData(group.Name, command + " - " + group.Name),
		)

		if key % 3 == 0 {
			buttons = append(buttons, line)

			line = nil
		}
	}

	return buttons
}

func HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if shouldNotAct(update) {
		// TODO: Log this why we don't act upon it
		return
	}

	// Override message with the reply
	if update.CallbackQuery != nil {
		handleReply(update, bot)
		return
	}

	// Create default message
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"I didn't understand...",
	)

	switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "off":
			msg.Text = "Which group of light do you want to turn off?"

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				listGroup("off")...
			)

		case "on":
			msg.Text = "Which group of light do you want to turn on?"

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				listGroup("on")...
			)
		default:
			msg.Text = "I don't know that command"
	}

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}