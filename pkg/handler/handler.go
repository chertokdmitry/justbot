package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/chertokdmitry/justbot/pkg/message"
	service "github.com/chertokdmitry/justbot/pkg/services"
	"github.com/chertokdmitry/justbot/pkg/telegram"
)

type Handler struct {
	services *service.Service
	update tgbotapi.Update
}

func NewHandler(services *service.Service, update tgbotapi.Update) *Handler {
	return &Handler{
		services: services,
		update: update,
	}
}

// Start main function to start app and route calls from bot
func Start() {
	update := telegram.GetUpdate()
	services := service.NewService()
	handlers := NewHandler(services, update)

	if update.CallbackQuery != nil {
		handlers.InlineRoute()
	} else {
		if update.Message.IsCommand() {
			if update.Message.Command() == message.StartCommand {
				handlers.CommandRoute()
			}
		}

		handlers.KeyboardRoute()
	}
}
