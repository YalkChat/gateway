package client

import (
	"context"
	"log"
	"time"

	"nhooyr.io/websocket"
)

type clientImpl struct {
	id      string
	conn    *websocket.Conn
	ctx     context.Context
	timeout time.Duration
}

// TODO: change hardcoded timeout
func NewClient(id string, conn *websocket.Conn) Client {
	return &clientImpl{
		id:      id,
		conn:    conn,
		timeout: time.Second * 5,
	}
}

func (c *clientImpl) ID() string {
	// Return the unique identifier for the client
	return c.id
}

func (c *clientImpl) SendMessage(messageType websocket.MessageType, p []byte) error {
	log.Printf("Sending message to client %s", c.id)
	return c.conn.Write(c.ctx, messageType, p)
}

func (c *clientImpl) ReadMessage() (messageType websocket.MessageType, p []byte, err error) {
	log.Printf("Reading message from client %s", c.id)
	return c.conn.Read(c.ctx)
}

func (c *clientImpl) SendMessageWithTimeout(messageType websocket.MessageType, data []byte) error {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
	defer cancel()

	// Write the message using the new context
	return c.conn.Write(ctx, messageType, data)
}

// Other methods as needed, such as receiving messages, handling events, etc.
