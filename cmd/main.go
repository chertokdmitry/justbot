package main

import (
	"github.com/chertokdmitry/justbot/pkg/handler"
	"github.com/chertokdmitry/justbot/pkg/telegram"
)

func main() {
	for update := range telegram.GetUpdatesChannel() {
		telegram.SetUpdate(update)

		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		handler.Start()
	}
}
