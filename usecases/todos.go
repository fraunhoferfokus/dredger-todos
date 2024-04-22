package usecases

import (
	"todos/entities"
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

func DeleteTodo(session string, id int) {
	Todos(session).Delete(string(id))
}

func DoneTodo(session string, id int) {
	Todos(session).ToggleDone(string(id))
}
