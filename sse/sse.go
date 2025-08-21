package sse

import (
	ssev2 "github.com/r3labs/sse/v2"
)

var SseServer *ssev2.Server

func init() {
	SseServer = ssev2.New()
}
