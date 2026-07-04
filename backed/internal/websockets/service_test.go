package websockets

import (
	"testing"
)

func TestBroadcastMessage(t *testing.T) {
	hub := NewHub()
	service := NewService(hub)

	go hub.Run()

	msg := Message{
		Type:    "chat",
		Content: "hello",
	}

	service.BroadcastMessage(1, msg)
}
