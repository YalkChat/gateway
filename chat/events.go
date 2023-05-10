package chat

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
	"yalk-backend/logger"

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

type EventChannels struct {
	Msg    chan *EventMessage
	Dm     chan *EventMessage
	Notify chan *EventMessage
	Cmd    chan *EventMessage
	Login  chan *EventMessage
	Logout chan *EventMessage
}

type EventMessage struct {
	Success   bool     `json:"success"`
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Sender    string   `json:"sender"`
	Receivers []string `json:"receivers"`
	Content   string   `json:"content,omitempty"`
}

type EventContext struct {
	NotifyChannel chan bool
	WaitGroup     *sync.WaitGroup
	PingTicket    *time.Ticker
	Connection    *websocket.Conn
	Request       *http.Request
}

func encodeEventMessage(event *EventMessage) ([]byte, error) {
	eventMessageJson, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return eventMessageJson, nil
}

func decodeEventMessage(jsonEvent []byte) (*EventMessage, error) {
	var event *EventMessage

	if err := json.Unmarshal(jsonEvent, &event); err != nil {
		logger.Err("BROAD", "Listener - Error deserializing JSON")
		return nil, err
	}
	return event, nil
}
