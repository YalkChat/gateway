package client

import (
	"context"

	"nhooyr.io/websocket"
)

type Client interface {
	// ID returns the unique identifier for the client
	ID() string

	SendMessage(ctx context.Context, messageType websocket.MessageType, p []byte) error
	ReadMessage(ctx context.Context) (messageType websocket.MessageType, p []byte, err error)

	// Other methods as needed, such as receiving messages, handling events, etc.
}
