package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/chertokdmitry/justbot/pkg/keyboard"
	"github.com/chertokdmitry/justbot/pkg/telegram"
)


func (h *Handler) CommandRoute() {
	msg := tgbotapi.NewMessage(h.update.Message.Chat.ID, h.update.Message.Text)
	msg.ReplyMarkup = keyboard.MainPage

	if _, err := telegram.GetBot().Send(msg); err != nil {
		logrus.Fatalf("error occured while send messasge to bot: %s", err.Error())
	}

	return
}