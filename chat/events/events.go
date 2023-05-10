package events

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
	"yalk/logger"

	"nhooyr.io/websocket"
)

// TODO: Return &ServerMessageChannels
func MakeEventChannels() *EventChannels {
	return &EventChannels{
		Msg:    make(chan *Event, 1),
		Dm:     make(chan *Event, 1),
		Notify: make(chan *Event, 1),
		Cmd:    make(chan *Event),
		Login:  make(chan *Event),
		Logout: make(chan *Event),
	}
}

type EventChannels struct {
	Msg    chan *Event
	Dm     chan *Event
	Notify chan *Event
	Cmd    chan *Event
	Login  chan *Event
	Logout chan *Event
}

type Event struct {
	Type string `gorm:"-" json:"type"`
}

type EventContext struct {
	NotifyChannel chan bool
	WaitGroup     *sync.WaitGroup
	PingTicket    *time.Ticker
	Connection    *websocket.Conn
	Request       *http.Request
}

func EncodeEventMessage(event *Event) ([]byte, error) {
	eventMessageJson, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return eventMessageJson, nil
}

func DecodeEventMessage(jsonEvent []byte) (*Event, error) {
	var event *Event

	if err := json.Unmarshal(jsonEvent, &event); err != nil {
		logger.Err("BROAD", "Listener - Error deserializing JSON")
		return nil, err
	}
	return event, nil
}
