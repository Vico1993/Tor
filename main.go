package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/heatxsink/go-hue/groups"
	"github.com/joho/godotenv"
)



func playWithLight() {
	gg := groups.New(os.Getenv("HUE_TEST_HOSTNAME"), os.Getenv("HUE_TEST_USERNAME"))
	allGroups, err := gg.GetAllGroups()
	if err != nil {
		fmt.Println("groups.GetAllGroups() ERROR: ", err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println("Groups")
	fmt.Println("------")
	for _, g := range allGroups {
		fmt.Printf("DEBUG ID: %d Name: %s\n", g.ID, g.Name)
		for _, lll := range g.Lights {
			fmt.Println("\t", lll)
		}

		// time.Sleep(time.Second * time.Duration(10))
		// _, err = gg.SetGroupState(g.ID, previousState)
		// if err != nil {
		// 	fmt.Println("groups.SetGroupState() ERROR: ", err)
		// 	os.Exit(1)
		// }
	}
}

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println(os.Getenv("TELEGRAM_BOT_TOKEN"))

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
    if err != nil {
        panic(err)
    }

    bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
    // to make sure Telegram knows we've handled previous values and we don't
    // need them repeated.
    updateConfig := tgbotapi.NewUpdate(0)

    // Tell Telegram we should wait up to 30 seconds on each request for an
    // update. This way we can get information just as quickly as making many
    // frequent requests without having to send nearly as many.
    updateConfig.Timeout = 30

    // Start polling Telegram for updates.
    updates := bot.GetUpdatesChan(updateConfig)

    // Let's go through each update that we're getting from Telegram.
    for update := range updates {
        // Telegram can send many types of updates depending on what your Bot
        // is up to. We only want to look at messages for now, so we can
        // discard any other updates.
        if update.Message == nil {
            continue
        }

		// Only the group can interact with my bot
		if update.FromChat().Type != "group" {
			continue
		}

		// ignore any non-command Messages
		if !update.Message.IsCommand() {
            continue
        }

        // Now that we know we've gotten a new message, we can construct a
        // reply! We'll take the Chat ID and Text from the incoming message
        // and use it to create a new message.
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

        // We'll also say that this message is a reply to the previous message.
        // For any other specifications than Chat ID or Text, you'll need to
        // set fields on the `MessageConfig`.
        msg.ReplyToMessageID = update.Message.MessageID

		switch update.Message.Command() {
        case "help":
            msg.Text = "I understand /sayhi and /status."
        case "sayhi":
            msg.Text = "Hi :)"
        case "status":
            msg.Text = "I'm ok."
		case "test":
            msg.Text = "Okay let's TEST"
			playWithLight()
        default:
            msg.Text = "I don't know that command"
        }

        // Okay, we're sending our message off! We don't care about the message
        // we just sent, so we'll discard it.
        if _, err := bot.Send(msg); err != nil {
            // Note that panics are a bad way to handle errors. Telegram can
            // have service outages or network errors, you should retry sending
            // messages or more gracefully handle failures.
            panic(err)
        }
    }
}
