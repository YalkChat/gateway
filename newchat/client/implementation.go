package client

import (
	"context"
	"log"
	"yalk/newchat/connection"

	"nhooyr.io/websocket"
)

type clientImpl struct {
	id   string
	conn connection.Connection
	ctx  context.Context
}

func NewClient(id string, conn connection.Connection) Client {
	return &clientImpl{
		id:   id,
		conn: conn,
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

// Other methods as needed, such as receiving messages, handling events, etc.
