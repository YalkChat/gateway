package clients

import (
	"context"
	"time"

	"nhooyr.io/websocket"
)

type Client struct {
	ID        uint
	Msgs      chan []byte
	CloseSlow func()
}

func ClientWriteWithTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}
