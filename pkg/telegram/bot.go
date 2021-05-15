package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/chertokdmitry/justbot/pkg/env"
	"sync"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

var singleton *Bot
var doOnce sync.Once

// GetMyBot singleton for bot instance
func GetMyBot() *Bot {
	doOnce.Do(func() {
		bot, err := tgbotapi.NewBotAPI(env.TOKEN)
		if err != nil {
			fmt.Println(err)
		}
		singleton = &Bot{bot: bot}
	})
	return singleton
}

func (b *Bot) get() *tgbotapi.BotAPI {
	return b.bot
}

// GetBot returns instance of bot
func GetBot() *tgbotapi.BotAPI {
	myBotStruct := GetMyBot()
	bot := myBotStruct.get()
	return bot
}

// GetChatId returns current chat id
func GetChatId() int64 {
	update := GetUpdate()
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	} else {
		return update.Message.Chat.ID
	}
}
