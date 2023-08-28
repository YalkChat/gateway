package connection

import (
	"context"

	"nhooyr.io/websocket"
)

type Connection interface {
	Read(ctx context.Context) (messageType websocket.MessageType, p []byte, err error)
	Write(ctx context.Context, messageType websocket.MessageType, p []byte) error
	Close(code websocket.StatusCode, reason string) error
}
