package chat

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

type Event interface {
	Type() string
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
	SaveToDb() error
}

type RawEvent struct {
	Type   string          `gorm:"-" json:"type"`
	Action string          `gorm:"-" json:"action"`
	UserID uint            `gorm:"-" json:"userId"`
	Data   json.RawMessage `gorm:"-" json:"data"`
}

func (event *RawEvent) Serialize() ([]byte, error) {
	return json.Marshal(event)
}

func (event *RawEvent) Deserialize(jsonEvent []byte) error {
	return json.Unmarshal(jsonEvent, event)
}

// TODO: Return &ServerMessageChannels
func MakeEventChannels() *EventChannels {
	return &EventChannels{
		Messages: make(chan *Message, 1),
		Accounts: make(chan *RawEvent, 1),
		Users:    make(chan *RawEvent, 1),
		Notify:   make(chan *RawEvent, 1),
		Cmd:      make(chan *RawEvent),
		Login:    make(chan *RawEvent),
		Logout:   make(chan *RawEvent),
	}
}

type EventChannels struct {
	Messages chan *Message
	Accounts chan *RawEvent
	Users    chan *RawEvent
	Notify   chan *RawEvent
	Cmd      chan *RawEvent
	Login    chan *RawEvent
	Logout   chan *RawEvent
}

type EventContext struct {
	NotifyChannel chan bool
	WaitGroup     *sync.WaitGroup
	PingTicket    *time.Ticker
	Connection    *websocket.Conn
	Request       *http.Request
}
