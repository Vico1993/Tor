package hue

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BuildListGroup(command string) [][]tgbotapi.InlineKeyboardButton {
	groups, err := GetAllGroup()
	if err != nil {
		panic(err)
	}

	var line []tgbotapi.InlineKeyboardButton
	var buttons [][]tgbotapi.InlineKeyboardButton
	for key, group := range groups {
		// Exclude custom Groups... where it's not a group
		if (strings.Contains(group.Name, "Custom group")) {
			continue
		}


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