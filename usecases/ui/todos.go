package ui

import (
	"dredgerTodos/async/publishers"

	"dredgerTodos/core/log"
	"dredgerTodos/entities"
	"dredgerTodos/usecases"
)

// Todos for each session
type TodosMap = map[string]*entities.ToDos

var currentTodos TodosMap = TodosMap{}

func Todos(session string) *entities.ToDos {
	todos, ok := currentTodos[session]
	if !ok {
		todos = entities.NewToDos()
		currentTodos[session] = todos
	}
	return todos
}

func AddTodo(session string, task string) {
	log.Debug().Str("task", task).Msg("Add todo message")

	label, found := usecases.CurrentLabels.GetLabel(session)
	if found {
		publishers.PublishAddTodo(entities.MsgAddToDo{Task: task, Done: false, Label: label})
	}
}

func DeleteTodo(session string, id int) {
	log.Debug().Int("key", id).Msg("Delete todo")

	label, found := usecases.CurrentLabels.GetLabel(session)
	if found {
		publishers.PublishDeleteTodo(entities.MsgDeleteToDo{Id: id, Label: label})
	}
}

func DoneTodo(session string, id int) {
	log.Debug().Int("key", id).Msg("Done todo")

	label, found := usecases.CurrentLabels.GetLabel(session)
	if found {
		publishers.PublishDoneTodo(entities.MsgDoneToDo{Id: id, Label: label})
	}
}
