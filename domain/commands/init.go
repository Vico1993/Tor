package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// type CmdParams struct {}

type Cmd interface {
	// Hack found to be able to get the command name
	Command() string
	Exec(Bot *tgbotapi.BotAPI, ChatId int64, params string) error
}

type baseCmd struct {
	Command	string
}

// Return the list of Command
var CmdList []Cmd

// Initialise the Command
// For now keep it hardcoded... find a better way later
func InitCmd() {
	// off
	CmdList = append(CmdList, &offCmd{
		baseCmd{
			Command: "off",
		},
	}, &onCmd{
		baseCmd{
			Command: "on",
		},
	}, &helpCmd{
		baseCmd{
			Command: "help",
		},
	})
}