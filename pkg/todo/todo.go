package todo


type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	ChatId		int	   `json:"chat_id"`
}

type ChatsList struct {
	Id     int
	ChatId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

