package main

import (
	"database/sql"
	"sync"

	"golang.org/x/time/rate"
)

type chatContext struct {
	// sessions sessions.SessionManager
	// db       *sql.DB
	// gorm     *gorm.DB
}

type chatServerConf struct {
	ServerID       string `json:"server_id"`
	DefaultChannel string `json:"default_channel"`
	TestKey        string `json:"test_key"`
	ConnType       string `json:"conn_type"`
}

type networkConf struct {
	URL  string
	IP   string
	Port string
}
type payload struct {
	Success bool   `json:"success"`
	Origin  string `json:"origin,omitempty"`
	Event   string `json:"event"`
	Data    any    `json:"data,omitempty"`
}

type channels struct {
	Msg     chan payload
	Dm      chan map[string]any
	Notify  chan payload
	Cmd     chan payload
	Conn    chan payload
	Disconn chan payload
}

type webSocketServer struct {
	SubscriberMessageBuffer int
	PublishLimiter          *rate.Limiter
	SubscribersMu           sync.Mutex
	channels                channels
	Clients                 map[string]*websocketClient
	DBconn                  *sql.DB
}

type websocketClient struct {
	Msgs      chan []byte
	CloseSlow func()
}
