package entities

import "strconv"

func NewToDos() *ToDos {
	return &ToDos{ToDos: map[string]ToDo{}, NextId: 0}
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

func (self *ToDos) Iter() map[string]ToDo {
	return self.ToDos
}
