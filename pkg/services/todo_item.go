package service

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/chertokdmitry/justbot/api"
	"github.com/chertokdmitry/justbot/pkg/env"
	"github.com/chertokdmitry/justbot/pkg/keyboard"
	"github.com/chertokdmitry/justbot/pkg/message"
	"github.com/chertokdmitry/justbot/pkg/telegram"
	"github.com/chertokdmitry/justbot/pkg/todo"
	"io/ioutil"
	"strconv"
)

var isItemTitle bool
var listId int

type TodoItemService struct {
}

func NewTodoItemService() *TodoItemService {
	return &TodoItemService{}
}

func (s *TodoItemService) SetItemTitleFalse() {
	isItemTitle = false
}

func (s *TodoItemService) SetListId(id int) {
	listId = id
}

func (s *TodoItemService) GetListId() int {
	return listId
}

func (s *TodoItemService) SetItemTitleTrue() {
	isItemTitle = true
}

func (s *TodoItemService) GetItemIsTitle() bool {
	return isItemTitle
}

// RequestItemTitle ask for title for task
func (s *TodoItemService) RequestItemTitle() {
	message.Send(message.EnterItemTitle)
	s.SetItemTitleTrue()
}

// Create task
func (s *TodoItemService) Create(tgMessage *tgbotapi.Message) {
	item := &ItemCreateRequest{int(telegram.GetChatId()), tgMessage.Text, false}
	jsonReq, err := json.Marshal(item)

	if err !=nil {
		logrus.Fatalf("error occured when marshal data: %s", err.Error())
	}

	request := api.ApiRequest{Url: env.API_HOST + "api/lists/" + strconv.Itoa(listId) + "/items", Method: "POST", Body: jsonReq}
	resp, err := request.NewApiRequest()

	if err != nil {
		logrus.Fatalf("error occured when get response from api: %s", err.Error())
	} else {
		var rows [][]tgbotapi.InlineKeyboardButton

		rows = append(rows, keyboard.MakeButtonRow(listId, message.AddItem, "new_item"))
		msgInline := tgbotapi.NewMessage(telegram.GetChatId(), message.ItemAdded)
		msgInline.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)

		if _, err := telegram.GetBot().Send(msgInline); err != nil {
			logrus.Fatalf("error occured when send message to bot: %s", err.Error())
		}

		defer resp.Body.Close()
	}

	s.SetItemTitleFalse()
	s.SetListId(0)
}

// GetListItems get tasks in list
func (s *TodoItemService) GetListItems(listItemsId int)  {
	var rows [][]tgbotapi.InlineKeyboardButton
	msgInline := tgbotapi.NewMessage(telegram.GetChatId(), message.ItemsInList)

	for _, item := range s.GetListItemsRequest(listItemsId) {
		var buttons []tgbotapi.InlineKeyboardButton
		buttons = append(buttons, keyboard.MakeButton(item.Id, item.Title, "item"))
		buttons = append(buttons, keyboard.MakeButton(item.Id, message.CheckItem,"delete_item" ))
		row := tgbotapi.NewInlineKeyboardRow(buttons...)
		rows = append(rows, row)
	}

	rows = append(rows, keyboard.MakeButtonRow(listItemsId, message.AddItem, "new_item"))
	msgInline.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)

	if _, err := telegram.GetBot().Send(msgInline); err != nil {
		logrus.Fatalf("error occured when send message to bot: %s", err.Error())
	}
}

// GetListItemsRequest request to api for tasks in list
func (s *TodoItemService) GetListItemsRequest(listId int) []todo.TodoItem {
	req := &GetListItemsRequest{int(telegram.GetChatId())}
	jsonReq, err := json.Marshal(req)

	if err !=nil {
		logrus.Fatalf("error occured when marshal data: %s", err.Error())
	}

	request := api.ApiRequest{Url: env.API_HOST + "api/lists/" + strconv.Itoa(listId) + "/items/all", Method: "POST", Body: jsonReq}
	resp, err := request.NewApiRequest()

	if err != nil {
		logrus.Fatalf("error occured when get response from api: %s", err.Error())
	}

	defer resp.Body.Close()
	dataJson, _ := ioutil.ReadAll(resp.Body)

	var getAllResp GetListItemsResponse

	if err := json.Unmarshal(dataJson, &getAllResp); err != nil {
		logrus.Fatalf("error occured when get unmarshal response from api: %s", err.Error())
	}

	return getAllResp.Data
}

// Delete make done task
func (s *TodoItemService) Delete(itemId int) {
	request := api.ApiRequest{Url: env.API_HOST + "api/items/" + strconv.Itoa(itemId) , Method: "DELETE"}
	_, err := request.NewApiRequest()

	if err != nil {
		logrus.Fatalf("error occured when get response from api: %s", err.Error())
	}

	message.Send(message.ItemClosed)
}
