package client

import (
	"nhooyr.io/websocket"
)

type Client interface {
	// ID returns the unique identifier for the client
	ID() string

	SendMessage(messageType websocket.MessageType, p []byte) error
	ReadMessage() (messageType websocket.MessageType, p []byte, err error)
	// Other methods as needed, such as receiving messages, handling events, etc.
}
