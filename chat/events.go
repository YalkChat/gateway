package chat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"yalk/logger"

	"nhooyr.io/websocket"
)

type RawEvent struct {
	Type string          `gorm:"-" json:"type"`
	Data json.RawMessage `gorm:"-" json:"Data"`
}

type Event interface {
	Type() string
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
}

// TODO: Return &ServerMessageChannels
func MakeEventChannels() *EventChannels {
	return &EventChannels{
		Msg:    make(chan *Message, 1),
		Dm:     make(chan *RawEvent, 1),
		Notify: make(chan *RawEvent, 1),
		Cmd:    make(chan *RawEvent),
		Login:  make(chan *RawEvent),
		Logout: make(chan *RawEvent),
	}
}

type EventChannels struct {
	Msg    chan *Message
	Dm     chan *RawEvent
	Notify chan *RawEvent
	Cmd    chan *RawEvent
	Login  chan *RawEvent
	Logout chan *RawEvent
}

type EventContext struct {
	NotifyChannel chan bool
	WaitGroup     *sync.WaitGroup
	PingTicket    *time.Ticker
	Connection    *websocket.Conn
	Request       *http.Request
}

func EncodeEventMessage(event *RawEvent) ([]byte, error) {
	eventMessageJson, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return eventMessageJson, nil
}

func DecodeEventMessage(jsonEvent []byte) (*RawEvent, error) {
	var event *RawEvent

	if err := json.Unmarshal(jsonEvent, &event); err != nil {
		logger.Err("EVENT", fmt.Sprintf("Error deserializing message to RawEvent: %v", err))
		return nil, err
	}
	return event, nil
}
