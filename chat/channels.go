package chat

import (
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

// TODO: Return &ServerMessageChannels
func makeEventChannels() *EventChannels {
	return &EventChannels{
		Msg:     make(chan Payload, 1),
		Dm:      make(chan map[string]any, 1),
		Notify:  make(chan Payload, 1),
		Cmd:     make(chan Payload),
		Conn:    make(chan Payload),
		Disconn: make(chan Payload),
	}
}

type EventChannels struct {
	Msg     chan Payload
	Dm      chan map[string]any
	Notify  chan Payload
	Cmd     chan Payload
	Conn    chan Payload
	Disconn chan Payload
	Test    chan string
}

type EventContext struct {
	Client        *Client
	Server        *Server
	NotifyChannel chan bool
	WaitGroup     *sync.WaitGroup
	PingTicket    *time.Ticker
	Connection    *websocket.Conn
	Request       *http.Request
}
