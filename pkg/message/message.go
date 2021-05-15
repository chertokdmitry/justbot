package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/chertokdmitry/justbot/pkg/telegram"
)

var botMessage string

// Get returns message
func Get(updateText string) string {
	if botMessage != "" {
		answer := botMessage
		botMessage = ""

		return answer
	} else {
		return updateText
	}
}

// Set save message to var
func Set(message string) {
	botMessage = message
}

// Send make message and send it to bot
func Send(message string) {
	bot := telegram.GetBot()
	msgMain := tgbotapi.NewMessage(telegram.GetChatId(), message)

	if _, err := bot.Send(msgMain); err != nil {
		logrus.Fatalf("error occured while sending message: %s", err.Error())
	}
}