package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/chertokdmitry/justbot/pkg/keyboard"
	"strconv"
)

// InlineRoute route calls for inline buttons
func (h *Handler) InlineRoute() {
	var data keyboard.Request

	if err := json.Unmarshal([]byte(h.update.CallbackQuery.Data), &data); err != nil {
		logrus.Fatalf("error occured when unmarshal data: %s", err.Error())
	}

	id, _ := strconv.Atoi(data.Id)

	switch data.Action {
	case "list_items":
		h.services.TodoItem.GetListItems(id)
		return

	case "new_item":
		h.services.TodoItem.SetItemTitleFalse()
		h.services.TodoItem.SetListId(id)
		h.services.TodoItem.RequestItemTitle()
		return

	case "delete_item":
		h.services.TodoItem.Delete(id)
		return

	case "delete_list":
		h.services.TodoList.Delete(id)
		return
	}
}