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

var isTitle bool

type TodoListService struct {
}

func NewTodoListService() *TodoListService {
	return &TodoListService{}
}

func (s *TodoListService) SetTitleFalse() {
	isTitle = false
}

func (s *TodoListService) SetTitleTrue() {
	isTitle = true
}

func (s *TodoListService) GetIsTitle() bool {
	return isTitle
}

func (s *TodoListService) GetTitle() {
	message.Send(message.EnterListTitle)
	s.SetTitleTrue()
}

func (s *TodoListService) Create(tgMessage *tgbotapi.Message) {
	todoList := &todo.TodoList{Title: tgMessage.Text, ChatId: int(telegram.GetChatId())}
	jsonReq, err := json.Marshal(todoList)

	if err !=nil {
		logrus.Fatalf("error occured when marshal data: %s", err.Error())
	}

	request := api.ApiRequest{Url: env.API_HOST + "api/lists", Method: "POST", Body: jsonReq}

	resp, err := request.NewApiRequest()
	if err != nil {
		logrus.Fatalf("error occured when get response from api: %s", err.Error())
	} else {
		defer resp.Body.Close()
	}
	message.Send(message.ListAdded)
	s.SetTitleFalse()
}

// GetAll get all lists
func (s *TodoListService) GetAll() {
	var rows [][]tgbotapi.InlineKeyboardButton
	update := telegram.GetUpdate()
	bot := telegram.GetBot()
	msgInline := tgbotapi.NewMessage(telegram.GetChatId(), update.Message.Text)

	if len(s.GetAllRequest()) > 0 {
		var wg sync.WaitGroup
		for _, list := range s.GetAllRequest() {
			wg.Add(1)
			rows = append(rows, s.GetListRow(list, &wg))
		}

		wg.Wait()

		msgInline.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
		if _, err := bot.Send(msgInline); err != nil {
			logrus.Fatalf("error occured when send message to bot: %s", err.Error())
		}
	} else {
		message.Send(message.EmptyLists)
	}
}

func (s *TodoListService) GetListRow(list todo.TodoList, wg *sync.WaitGroup) []tgbotapi.InlineKeyboardButton {
	defer wg.Done()
	var buttons []tgbotapi.InlineKeyboardButton
	buttons = append(buttons, keyboard.MakeButton(list.Id, list.Title,"list_items" ))
	buttons = append(buttons, keyboard.MakeButton(list.Id, message.DeleteList,"delete_list" ))

	return tgbotapi.NewInlineKeyboardRow(buttons...)
}

// GetAllRequest request to get all lists from api
func (s *TodoListService) GetAllRequest() []todo.TodoList {
	req := &GetAllRequest{telegram.GetChatId()}
	jsonReq, err := json.Marshal(req)

	if err !=nil {
		logrus.Fatalf("error occured when marshal data: %s", err.Error())
	}
	request := api.ApiRequest{Url: env.API_HOST + "api/lists/all", Method: "POST", Body: jsonReq}

	resp, err := request.NewApiRequest()
	if err != nil {
		logrus.Fatalf("error occured when get response from api: %s", err.Error())
	}

	defer resp.Body.Close()
	dataJson, _ := ioutil.ReadAll(resp.Body)

	var getAllResp GetAllResponse

	if err := json.Unmarshal(dataJson, &getAllResp); err != nil {
		logrus.Fatalf("error occured when get unmarshal response from api: %s", err.Error())
	}

	return getAllResp.Data
}

// GetById get list by by id
func (s *TodoListService) GetById(chatId, listId int) (todo.TodoList, error) {

	return todo.TodoList{}, nil
}

// Delete make list inactive
func (s *TodoListService) Delete(listId int) {
	request := api.ApiRequest{Url: env.API_HOST + "api/lists/" + strconv.Itoa(listId) , Method: "DELETE"}
	_, err := request.NewApiRequest()

	if err != nil {
		logrus.Fatalf("error occured when get response from api: %s", err.Error())
	}

	message.Send(message.ListDeleted)
}

func (s *TodoListService) Update(chatId, listId int, input todo.UpdateListInput) error {
	return nil
}
