package handlers

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

// TODO: Check test... not working very well
func TestShouldReactShouldReturnFalseIfUpdateNotMessage(t *testing.T) {
	update := tgbotapi.Update{
		UpdateID: 1,
		Message: nil,
		EditedMessage: &tgbotapi.Message{
			MessageID: 1,
			Chat: &tgbotapi.Chat{
				ID: 1,
				Type: "group",
			},
		},
	}

	assert.Equal(t, false, shouldNotReact(update))
}

func TestShouldReactShouldReturnFalseIfNotGroupMessage(t *testing.T) {
	update := tgbotapi.Update{
		UpdateID: 1,
		Message: &tgbotapi.Message{
			MessageID: 1,
			Text: "Hello",
			Chat: &tgbotapi.Chat{
				ID: 1,
				Type: "private",
			},
		},
	}

	assert.Equal(t, false, shouldNotReact(update))
}

func TestShouldReactShouldReturnFalseIfEmptyEntity(t *testing.T) {
	update := tgbotapi.Update{
		UpdateID: 1,
		Message: &tgbotapi.Message{
			MessageID: 1,
			Text: "Hello",
			Chat: &tgbotapi.Chat{
				ID: 1,
				Type: "group",
			},
			Entities: make([]tgbotapi.MessageEntity, 0),
		},
	}

	assert.Equal(t, false, shouldNotReact(update))
}

func TestShouldReactShouldReturnFalseIfNotGoodTypeOfEntityReceived(t *testing.T) {
	update := tgbotapi.Update{
		UpdateID: 1,
		Message: &tgbotapi.Message{
			MessageID: 1,
			Text: "Hello",
			Chat: &tgbotapi.Chat{
				ID: 1,
				Type: "group",
			},
			Entities: []tgbotapi.MessageEntity{
				{
					Type: "hashtag",
				},
			},
		},
	}

	assert.Equal(t, false, shouldNotReact(update))
}

// func TestShouldReactShouldReturnTrueIfCorrectMessage(t *testing.T) {
// 	update := tgbotapi.Update{
// 		UpdateID: 1,
// 		Message: &tgbotapi.Message{
// 			MessageID: 1,
// 			Text: "Hello",
// 			Chat: &tgbotapi.Chat{
// 				ID: 1,
// 				Type: "group",
// 			},
// 			Entities: []tgbotapi.MessageEntity{
// 				{
// 					Type: "bot_command",
// 				},
// 			},
// 		},
// 	}

// 	assert.Equal(t, false, shouldReact(update))
// }