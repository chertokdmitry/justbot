package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/chertokdmitry/justbot/pkg/keyboard"
	"github.com/chertokdmitry/justbot/pkg/message"
	"github.com/chertokdmitry/justbot/pkg/telegram"
)

// KeyboardRoute route calls for keyboard buttons and message text
func (h *Handler) KeyboardRoute() {
	if h.update.Message.Text == "/" {
		msg := tgbotapi.NewMessage(h.update.Message.Chat.ID, h.update.Message.Text)
		msg.ReplyMarkup = keyboard.MainPage

		if _, err := telegram.GetBot().Send(msg); err != nil {
			logrus.Fatalf("error occured while send messasge to bot: %s", err.Error())
		}
	}

	switch h.update.Message.Text {
	case message.NewList:
		h.services.TodoList.SetTitleFalse()
		h.services.TodoList.GetTitle()
		return

	case message.AllLists:
		h.services.TodoList.GetAll()
		return
	}

	if h.services.TodoList.GetIsTitle() == true && h.update.Message.Text != "" {
		h.services.TodoList.Create(h.update.Message)
		return
	}

	if h.services.TodoItem.GetItemIsTitle() == true && h.update.Message.Text != "" {
		h.services.TodoItem.Create(h.update.Message)
		return
	}
}
