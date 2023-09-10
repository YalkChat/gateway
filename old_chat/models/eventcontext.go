package models

import (
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

type EventContext struct {
	NotifyChannel chan bool
	WaitGroup     *sync.WaitGroup
	PingTicket    *time.Ticker
	Connection    *websocket.Conn
	Request       *http.Request
}
