package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
)

type Update struct {
	val tgbotapi.Update
}

var singletonUpdate *Update
var doOnceUpdate sync.Once

// GetMyUpdate singleton for update
func GetMyUpdate() *Update {
	doOnceUpdate.Do(func() {
		singletonUpdate = &Update{}
	})
	return singletonUpdate
}

func (u *Update) getVal() tgbotapi.Update {
	return u.val
}

// GetUpdate returns current update
func GetUpdate() tgbotapi.Update {
	myUpdateStruct := GetMyUpdate()
	update := myUpdateStruct.getVal()
	return update
}

// SetUpdate save current update to struct
func SetUpdate(update tgbotapi.Update) {
	myUpdateStruct := GetMyUpdate()
	myUpdateStruct.val = update
}

// GetUpdatesChannel returns channel to get updates
func GetUpdatesChannel() tgbotapi.UpdatesChannel {
	bot := GetBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	tgUpdates, _ := bot.GetUpdatesChan(u)

	return tgUpdates
}

