package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/chertokdmitry/justbot/pkg/todo"
	"sync"
)

type TodoList interface {
	Create(message *tgbotapi.Message)
	GetAll()
	GetById(chatId, listId int) (todo.TodoList, error)
	Update(chatId, listId int, input todo.UpdateListInput) error
	GetIsTitle() bool
	GetTitle()
	SetTitleFalse()
	SetTitleTrue()
	GetAllRequest() []todo.TodoList
	Delete(listId int)
	GetListRow(list todo.TodoList, wg *sync.WaitGroup) []tgbotapi.InlineKeyboardButton
}

type TodoItem interface {
	Create(message *tgbotapi.Message)
	GetListItems(listId int)
	GetListItemsRequest(listId int) []todo.TodoItem
	SetItemTitleFalse()
	SetItemTitleTrue()
	RequestItemTitle()
	SetListId(id int)
	GetListId() int
	GetItemIsTitle() bool
	Delete(itemId int)
	GetItemRow(item todo.TodoItem, wg *sync.WaitGroup) []tgbotapi.InlineKeyboardButton
}

type Service struct {
	TodoList
	TodoItem
}

func NewService() *Service {
	return &Service{
		TodoList:      NewTodoListService(),
		TodoItem:      NewTodoItemService(),
	}
}