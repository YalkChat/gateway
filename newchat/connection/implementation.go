package connection

import (
	"context"
	"log"

	"nhooyr.io/websocket"
)

type connectionImpl struct {
	conn *websocket.Conn
}

func NewConnection(conn *websocket.Conn) Connection {
	return &connectionImpl{
		conn: conn,
	}
}

func (c *connectionImpl) Read(ctx context.Context) (messageType websocket.MessageType, p []byte, err error) {
	return c.conn.Read(ctx)
}

func (c *connectionImpl) Write(ctx context.Context, messageType websocket.MessageType, p []byte) error {
	return c.conn.Write(ctx, messageType, p)
}

func (c *connectionImpl) Close(code websocket.StatusCode, reason string) error {
	log.Println("Closing WebSocket connection")
	return c.conn.Close(code, reason)
}
