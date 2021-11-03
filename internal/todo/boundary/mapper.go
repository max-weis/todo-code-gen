package boundary

import (
	"todo-code-gen/internal/todo/entity"
)

func todosToList(todos []entity.ListTodosRow) *[]TodoList {
	var list []TodoList

	for _, todo := range todos {
		list = append(list, toListDTO(entity.Todo{
			ID:    todo.ID,
			Title: todo.Title,
		}))
	}

	return &list
}

func toListDTO(todo entity.Todo) TodoList {
	list := TodoList{
		Done:  stateToBool(todo.Status),
		Id:    int(todo.ID),
		Title: todo.Title,
	}

	return list
}

func toFullDTO(todo entity.Todo) *TodoFull {
	full := TodoFull{
		TodoBase: toBaseDTO(todo),
		Id:       int(todo.ID),
	}

	return &full
}

func toBaseDTO(todo entity.Todo) TodoBase {
	base := TodoBase{
		Description: &todo.Description.String,
		Done:        stateToBool(todo.Status),
		Title:       todo.Title,
	}

	return base
}

func stateToBool(state int32) bool {
	if state <= 0 {
		return false
	}

	return true
}
