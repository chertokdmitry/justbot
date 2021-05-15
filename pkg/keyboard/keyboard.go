package keyboard

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/chertokdmitry/justbot/pkg/message"
	"strconv"
)

var MainPage = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(message.NewList),
		tgbotapi.NewKeyboardButton(message.AllLists),
	),
)

type Request struct {
	Action string `json:"action"`
	Id     string `json:"id"`
}

// MakeButton returns tgbotapi button
func MakeButton(id int, title, action string) tgbotapi.InlineKeyboardButton {
	request := &Request{
		Action: action,
		Id:     strconv.Itoa(id),
	}

	var msg, _ = json.Marshal(request)

	return  tgbotapi.NewInlineKeyboardButtonData(title, string(msg))
}

// MakeButtonRow returns tgbotapi row of buttons
func MakeButtonRow(id int, title, action string) []tgbotapi.InlineKeyboardButton {
	newItemButton := MakeButton(id, title, action)

	return tgbotapi.NewInlineKeyboardRow(newItemButton)
}