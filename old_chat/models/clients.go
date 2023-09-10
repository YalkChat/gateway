package models

import (
	"context"
	"log"
	"time"

	"nhooyr.io/websocket"
)

type Client struct {
	ID        uint
	Msgs      chan []byte
	CloseSlow func()
}

// TODO: Method on server Type
func ClientWriteWithTimeout(id uint, ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	log.Printf("Sending with Timeout to %d", id)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}
