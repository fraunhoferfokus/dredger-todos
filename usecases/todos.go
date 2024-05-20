package usecases

import (
	"dredgerTodos/entities"
	"strconv"
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
	Todos(session).Delete(strconv.Itoa(id))
}

func DoneTodo(session string, id int) {
	Todos(session).ToggleDone(strconv.Itoa(id))
}
