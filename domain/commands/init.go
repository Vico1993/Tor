package commands

import (
	"github.com/Vico1993/Tor/domain/hue"
	"github.com/Vico1993/Tor/domain/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
	CmdList = []Cmd{
		&offCmd{
			baseCmd{
				Command: "off",
			},
		},
		&onCmd{
			baseCmd{
				Command: "on",
			},
		}, &helpCmd{
			baseCmd{
				Command: "help",
			},
		},
	}

	// Build dynamic command
	groups, err := hue.GetAllGroup()
	if err != nil {
		panic(err)
	}

	for _, grp := range groups {
		CmdList = append(CmdList, &groupCmd{
			group: grp,
			baseCmd: baseCmd{
				Command: util.ToCamelCase(grp.Name),
			},
		})
	}
}