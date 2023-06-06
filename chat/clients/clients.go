package clients

import (
	"context"
	"fmt"
	"time"
	"yalk/logger"

	"nhooyr.io/websocket"
)

type Client struct {
	ID        uint
	Msgs      chan []byte
	CloseSlow func()
}

func ClientWriteWithTimeout(id uint, ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	logger.Info("CLNT", fmt.Sprintf("Sending with Timeout to %d", id))
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}
