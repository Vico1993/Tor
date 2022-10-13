package commands

import (
	"github.com/Vico1993/Tor/domain/hue"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type offCmd struct {
	baseCmd
}

// Off command
func (cmd *offCmd) Exec(bot *tgbotapi.BotAPI, chatId int64, params string) error {
	msg := tgbotapi.NewMessage(
		chatId,
		"Which group of light do you want to turn off?",
	)

	if params == "" {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			hue.BuildListGroup("off")...
		)
	} else {
		msg.Text = "Shutting down: " + params
		hue.ShutDownGroup(params)
	}

	if _, err := bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (cmd *offCmd) Command() string {
	return cmd.baseCmd.Command
}