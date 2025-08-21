// Edit this file, as it is a specific handler function for your service
package publishers

import (
	"dredgerTodos/async"
	"encoding/json"
	"fmt"
	//noch an core logger anpassen
)

func PublishDeleteTodo(v any) error {
	// next line depends on contentType
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("Marshalling error")
	}

	return async.NC.Publish("todos.deletetodo", data)
}
