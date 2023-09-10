package client

import (
	"nhooyr.io/websocket"
)

type Client interface {
	// ID returns the unique identifier for the client
	ID() uint

	SendMessage(websocket.MessageType, []byte) error
	SendMessageWithTimeout(websocket.MessageType, []byte) error
	ReadMessage() (websocket.MessageType, []byte, error)
	// Other methods as needed, such as receiving messages, handling events, etc.
}
