package main

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/time/rate"
	"nhooyr.io/websocket"
)

func newWebSocketServer(db *sql.DB, channels channels) *webSocketServer {
	wss := &webSocketServer{
		DBconn:                  db,
		SubscriberMessageBuffer: 16,
		Clients:                 make(map[string]*websocketClient),
		PublishLimiter:          rate.NewLimiter(rate.Every(time.Millisecond*100), 8),
	}

	return wss
}

// func (websock *webSocketServer) DeleteSubscriber(s *websocketClient, id string) {

// }

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}
