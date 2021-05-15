package service

import "github.com/chertokdmitry/justbot/pkg/todo"

type GetAllRequest struct {ChatId int64	`json:"chat_id"`}

type GetAllResponse struct {
	Data	[]todo.TodoList `json:"data"`
}
