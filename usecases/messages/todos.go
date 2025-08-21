package messages

import (
	"dredgerTodos/core/log"
	"dredgerTodos/entities"
	"dredgerTodos/sse"
	"dredgerTodos/usecases"

	ssev2 "github.com/r3labs/sse/v2"

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

func AddTodoMsg(label string, task string) {
	log.Debug().Str("task", task).Msg("Add todo message")

	for _, session := range usecases.CurrentLabels.GetSessions(label) {
		Todos(session).Add(task)
	}
	sse.SseServer.Publish("refresh", &ssev2.Event{
		Event: []byte("Refresh"),
		Data:  []byte("-"),
	})
}

func DeleteTodoMsg(label string, id int) {
	log.Debug().Int("key", id).Msg("Delete todo message")

	for _, session := range usecases.CurrentLabels.GetSessions(label) {
		Todos(session).Delete(strconv.Itoa(id))
	}
	sse.SseServer.Publish("refresh", &ssev2.Event{
		Event: []byte("Refresh"),
		Data:  []byte("-"),
	})
}

func DoneTodoMsg(label string, id int) {
	log.Debug().Int("key", id).Msg("Done todo message")

	for _, session := range usecases.CurrentLabels.GetSessions(label) {
		Todos(session).ToggleDone(strconv.Itoa(id))
	}
	sse.SseServer.Publish("refresh", &ssev2.Event{
		Event: []byte("Refresh"),
		Data:  []byte("-"),
	})
}
