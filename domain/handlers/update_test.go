package handlers

// TODO: Check test... not working very well
// func TestShouldActShouldReturnTrueIfUpdateNotMessage(t *testing.T) {
// 	update := tgbotapi.Update{
// 		UpdateID: 1,
// 		Message: nil,
// 		EditedMessage: &tgbotapi.Message{
// 			MessageID: 1,
// 			Chat: &tgbotapi.Chat{
// 				ID: 1,
// 				Type: "group",
// 			},
// 		},
// 	}

// 	assert.Equal(t, true, shouldNotAct(update))
// }

// func TestShouldActShouldReturnTrueIfNotGroupMessage(t *testing.T) {
// 	update := tgbotapi.Update{
// 		UpdateID: 1,
// 		Message: &tgbotapi.Message{
// 			MessageID: 1,
// 			Text: "Hello",
// 			Chat: &tgbotapi.Chat{
// 				ID: 1,
// 				Type: "private",
// 			},
// 		},
// 	}

// 	assert.Equal(t, true, shouldNotAct(update))
// }

// func TestShouldActShouldReturnTrueIfEmptyEntity(t *testing.T) {
// 	update := tgbotapi.Update{
// 		UpdateID: 1,
// 		Message: &tgbotapi.Message{
// 			MessageID: 1,
// 			Text: "Hello",
// 			Chat: &tgbotapi.Chat{
// 				ID: 1,
// 				Type: "group",
// 			},
// 			Entities: make([]tgbotapi.MessageEntity, 0),
// 		},
// 	}

// 	assert.Equal(t, true, shouldNotAct(update))
// }

// func TestShouldActShouldReturnTrueIfNotGoodTypeOfEntityReceived(t *testing.T) {
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
// 					Type: "hashtag",
// 				},
// 			},
// 		},
// 	}

// 	assert.Equal(t, true, shouldNotAct(update))
// }

// // func TestShouldActShouldReturnTrueIfCorrectMessage(t *testing.T) {
// // 	update := tgbotapi.Update{
// // 		UpdateID: 1,
// // 		Message: &tgbotapi.Message{
// // 			MessageID: 1,
// // 			Text: "Hello",
// // 			Chat: &tgbotapi.Chat{
// // 				ID: 1,
// // 				Type: "group",
// // 			},
// // 			Entities: []tgbotapi.MessageEntity{
// // 				{
// // 					Type: "bot_command",
// // 				},
// // 			},
// // 		},
// // 	}

// // 	assert.Equal(t, false, ShouldAct(update))
// // }