package entities

import (
	"maps"
	"slices"
	"strconv"
)

func NewToDos() *ToDos {
	return &ToDos{ToDos: map[string]ToDo{}, NextId: 0}
}

func (self *ToDos) Ids() []string {
	return slices.Sorted(maps.Keys(self.ToDos))
}

func (self *ToDos) Task(id string) string {
	return self.ToDos[id].Task
}

func (self *ToDos) Done(id string) bool {
	return self.ToDos[id].Done
}

func (self *ToDos) Add(task string) {
	self.ToDos[strconv.Itoa(self.NextId)] = ToDo{Done: false, Task: task}
	self.NextId += 1
}

func (self *ToDos) ToggleDone(id string) {
	t := self.ToDos[id]
	t.Done = !t.Done
	self.ToDos[id] = t
}

func (self *ToDos) Delete(id string) {
	delete(self.ToDos, id)
}
