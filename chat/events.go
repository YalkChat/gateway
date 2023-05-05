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
		Msg:    make(chan *EventMessage, 1),
		Dm:     make(chan *EventMessage, 1),
		Notify: make(chan *EventMessage, 1),
		Cmd:    make(chan *EventMessage),
		Login:  make(chan *EventMessage),
		Logout: make(chan *EventMessage),
	}
}

type EventMessage struct {
	Success   bool     `json:"success"`
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Sender    string   `json:"sender"`
	Receivers []string `json:"receivers"`
	Message   string   `json:"message,omitempty"`
}

type EventChannels struct {
	Msg    chan *EventMessage
	Dm     chan *EventMessage
	Notify chan *EventMessage
	Cmd    chan *EventMessage
	Login  chan *EventMessage
	Logout chan *EventMessage
}

type EventContext struct {
	NotifyChannel chan bool
	WaitGroup     *sync.WaitGroup
	PingTicket    *time.Ticker
	Connection    *websocket.Conn
	Request       *http.Request
}
