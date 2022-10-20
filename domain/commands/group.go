package commands

import (
	"fmt"
	"strings"

	"github.com/Vico1993/Tor/domain/hue"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/heatxsink/go-hue/groups"
)

type groupCmd struct {
	baseCmd
	group groups.Group
}

func (cmd *groupCmd) Exec(bot *tgbotapi.BotAPI, chatId int64, params string) error {
	if params == "" {
		// Return help command for this group
		return nil
	}

	sliceParams := strings.Split(params, " ")

	if sliceParams[0] == "" {
		return nil
	}

	// First params need to be scene name or on/off
	if sliceParams[0] == "on" {
		hue.PowerUpGroup(cmd.group.Name)
		return nil
	}

	if sliceParams[0] == "off" {
		hue.ShutDownGroup(cmd.group.Name)
		return nil
	}

	// Check scene for the group

	// Second is either not set
	// If not set set it to 100% or 0

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
