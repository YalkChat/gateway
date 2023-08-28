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

func (c *clientImpl) SendMessage(ctx context.Context, messageType websocket.MessageType, p []byte) error {
	log.Printf("Sending message to client %s", c.id)
	return c.conn.Write(ctx, messageType, p)
}

func (c *clientImpl) ReadMessage(ctx context.Context) (messageType websocket.MessageType, p []byte, err error) {
	log.Printf("Reading message from client %s", c.id)
	return c.conn.Read(ctx)
}

// func (c *clientImpl) SendEvent(event event.Event) error {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	// Serialize the event and send it through the WebSocket connection
// 	// ...
// 	return nil
// }

// func (c *clientImpl) ReadEvent() (*event.Event, error) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	// Deserialize the incoming event for the WebSocket connection
// 	// ...
// 	return nil, nil
// }

// Other methods as needed, such as receiving messages, handling events, etc.
