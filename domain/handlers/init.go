package handlers

import (
	"github.com/Vico1993/Tor/domain/commands"
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

// TODO: Find a way to manage Error Gracefully because panicing is not a solution
func HandleEvent(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if shouldNotAct(update) {
		// TODO: Log this why we don't act upon it
		return
	}

	var ctx handlerContext
	// If Reply Event
	if update.CallbackQuery != nil {
		ctx = buildReplyContext(update, bot)
	} else {
		ctx = buildUpdateContext(update, bot)
	}

	commandHandled := false
	for _, cmd := range commands.CmdList {
		if cmd.Command() == ctx.Command {
			err := cmd.Exec(bot, ctx.ChatId, ctx.CommandParam)
			if err != nil {
				panic(err)
			}

			commandHandled = true
			break
		}
	}

	// If we couldn't find the command use asked by the user
	if !commandHandled {
		msg := tgbotapi.NewMessage(
			ctx.ChatId,
			"Couldn't find your command",
		)

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}