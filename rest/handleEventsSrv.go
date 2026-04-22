// Edit this file, as it is a specific for your service
package rest

import (
	"dredgerTodos/sse"
)

func handleEventsSvc() {
	sse.SseServer.CreateStream("refresh")
}
