package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type helpCmd struct {
	baseCmd
}

func buildHelpMessage() string {
	text := "Hey! You need help! This how I can help you today"

	for _, cmd := range CmdList {
		text += "\n /" + cmd.Command() + ": " + cmd.Description()
	}

	return text
}

// Help command
func (cmd *helpCmd) Exec(bot *tgbotapi.BotAPI, chatId int64, params string) error {
	msg := tgbotapi.NewMessage(
		chatId,
		buildHelpMessage(),
	)

	if _, err := bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (cmd *helpCmd) Command() string {
	return cmd.baseCmd.Command
}

func (cmd *helpCmd) Description() string {
	return cmd.baseCmd.Description
}
