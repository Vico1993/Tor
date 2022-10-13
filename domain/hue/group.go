package hue

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/lights"
)

// From a name detect if it's a group otherwise send nil
func IsAGroup(name string) *groups.Group {
	group, err := findCorrectGroup(name)
	if err != nil {
		panic(err)
	}

	return group
}

// Get a list of quick replies with the correct list of groups
func BuildListGroup(command string) [][]tgbotapi.InlineKeyboardButton {
	groups, err := getAllGroup()
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

	if line != nil {
		buttons = append(buttons, line)
	}

	return buttons
}

// Shutdown a all group
func ShutDownGroup(name string) {
	group, err := findCorrectGroup(name)
	if err != nil {
		panic(err)
	}

	fmt.Println("SHUTDOWN")

	_, err = getGroupClient().SetGroupState(
		group.ID,
		lights.State{
			On: false,
			Bri: 0,
			TransitionTime: 1,
		},
	)

	if err != nil {
		panic(err)
	}
}

// Turn on a all group
func PowerUpGroup(name string) {
	group, err := findCorrectGroup(name)
	if err != nil {
		panic(err)
	}

	_, err = getGroupClient().SetGroupState(
		group.ID,
		lights.State{
			On: true,
			Bri: 200,
			TransitionTime: 1,
		},
	)

	if err != nil {
		panic(err)
	}
}
