package main

import (
	"log"
	"os"

	"github.com/Vico1993/Tor/domain/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// TODO: Handle Errors

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// SETUP BOT
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
    if err != nil {
        panic(err)
    }

    bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 30

    // Start polling Telegram for updates.
    updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		handlers.HandleUpdate(update, bot)
	}
}