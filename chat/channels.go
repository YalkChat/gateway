package chat

import (
	"net/http"
	"sync"
	"time"

	"gorm.io/gorm"
	"nhooyr.io/websocket"
)

// TODO: Return &ServerMessageChannels
func makeMessageChannels() MessageChannels {
	return MessageChannels{
		Msg:     make(chan Payload, 1),
		Dm:      make(chan map[string]any, 1),
		Notify:  make(chan Payload, 1),
		Cmd:     make(chan Payload),
		Conn:    make(chan Payload),
		Disconn: make(chan Payload),
	}
}

type MessageChannels struct {
	Msg     chan Payload
	Dm      chan map[string]any
	Notify  chan Payload
	Cmd     chan Payload
	Conn    chan Payload
	Disconn chan Payload
}

type MessageChannelsContext struct {
	NotifyChannel chan bool
	WaitGroup     *sync.WaitGroup
	PingTicket    *time.Ticker
	Connection    *websocket.Conn
	Request       *http.Request
	ClientData    *Client
	Channels      *MessageChannels
	Db            *gorm.DB
}
