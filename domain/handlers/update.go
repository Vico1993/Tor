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

	botRespond(botParameter{
		ChatId: update.Message.Chat.ID,
		Bot: bot,
		Command: update.Message.Command(),
		CommandParam: "",
	})
}

type botParameter struct {
	Bot 			*tgbotapi.BotAPI
	ChatId 			int64
	Command 		string
	CommandParam 	string
}

func botRespond(params botParameter) {
	// Create default message
	msg := tgbotapi.NewMessage(
		params.ChatId,
		"I didn't understand...",
	)

	switch params.Command {
		case "help":
			msg.Text = "I understand /off and /on."
		case "off":
			if params.CommandParam != "" {
				msg.Text = "Shutting down: " + params.CommandParam
				hue.ShutDownGroup(params.CommandParam)
			} else {
				msg.Text = "Which group of light do you want to turn off?"
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
					hue.BuildListGroup("off")...
				)
			}
		case "on":
			if params.CommandParam != "" {
				msg.Text = "Powering up: " + params.CommandParam
				hue.PowerUpGroup(params.CommandParam)
			} else {
				msg.Text = "Which group of light do you want to turn on?"
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
					hue.BuildListGroup("on")...
				)
			}
		default:
			msg.Text = "I don't know that command"
	}

	if _, err := params.Bot.Send(msg); err != nil {
		panic(err)
	}
}