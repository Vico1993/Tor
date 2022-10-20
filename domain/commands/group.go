package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/heatxsink/go-hue/groups"
)

type groupCmd struct {
	baseCmd
	group groups.Group
}

func (cmd *groupCmd) Exec(bot *tgbotapi.BotAPI, chatId int64, params string) error {
	fmt.Println("Group Command!")
	fmt.Println("Here are the parameters:", params)

	return nil
}

func (cmd *groupCmd) Command() string {
	return cmd.baseCmd.Command
}

func (cmd *groupCmd) Description() string {
	return cmd.baseCmd.Description
}
