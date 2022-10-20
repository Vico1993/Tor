package commands

import (
	"github.com/Vico1993/Tor/domain/hue"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type onCmd struct {
	baseCmd
}

// On command
func (cmd *onCmd) Exec(bot *tgbotapi.BotAPI, chatId int64, params string) error {
	msg := tgbotapi.NewMessage(
		chatId,
		"Which group of light do you want to turn on?",
	)

	if params == "" {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			hue.BuildListGroup("on")...,
		)
	} else {
		msg.Text = "Powering up: " + params
		hue.PowerUpGroup(params)
	}

	if _, err := bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (cmd *onCmd) Command() string {
	return cmd.baseCmd.Command
}

func (cmd *onCmd) Description() string {
	return cmd.baseCmd.Description
}
