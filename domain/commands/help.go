package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type helpCmd struct {
	baseCmd
}

// Help command
func (cmd *helpCmd) Exec(bot *tgbotapi.BotAPI, chatId int64, params string) error {
	msg := tgbotapi.NewMessage(
		chatId,
		"I understand /off and /on.",
	)

	if _, err := bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (cmd *helpCmd) Command() string {
	return cmd.baseCmd.Command
}