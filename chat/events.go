package chat

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
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
	UserID string          `gorm:"-" json:"userId"`
	Token  jwt.Token       `gorm:"-" json:"token"`
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
