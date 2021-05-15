package service

import "github.com/chertokdmitry/justbot/pkg/todo"

type ItemCreateRequest struct {
	ChatId      int    `json:"chat_id"`
	Title       string `json:"title"`
	Done        bool   `json:"done"`
}

type GetListItemsRequest struct {
	ChatId      int    `json:"chat_id"`
}

type GetListItemsResponse struct {
	Data	[]todo.TodoItem `json:"data"`
}
