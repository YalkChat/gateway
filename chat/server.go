package chat

import (
	"fmt"
	"sync"
	"time"

	"yalk-backend/logger"

	"golang.org/x/time/rate"
	"gorm.io/gorm"
	"nhooyr.io/websocket"
)

type ChatServer interface {
	RegisterClient(*websocket.Conn, string)
	SendMessage(*EventMessage)
	SendMessageToAll(*EventMessage)
	Sender(*Client, *EventContext)
	Receiver(*EventContext)
	HandlePayload([]byte)
}

// TODO: db
func NewServer(bufferLenght int, dbConfig *PgConf) *Server {

	sendLimiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 8)
	clientsMap := make(map[string]*Client)
	messageChannels := makeEventChannels()
	db, err := openDbConnection(dbConfig)

	if err != nil {
		logger.Err("SRV", fmt.Sprintf("Error opening db connection: %v\n", err))
		return nil
	}
	chatServer := &Server{
		SendLimiter:          sendLimiter,
		Clients:              clientsMap,
		ClientsMessageBuffer: bufferLenght,
		Channels:             messageChannels,
		Db:                   db,
	}

	return chatServer
}

type Server struct {
	SendLimiter          *rate.Limiter
	Clients              map[string]*Client
	ClientsMu            sync.Mutex
	ClientsMessageBuffer int
	Channels             *EventChannels
	Db                   *gorm.DB
}

func (server *Server) RegisterClient(conn *websocket.Conn, id string) *Client {
	messageChan := make(chan []byte, server.ClientsMessageBuffer)

	client := &Client{
		Msgs: messageChan,
		CloseSlow: func() {
			conn.Close(websocket.StatusPolicyViolation, "connection too slow to keep up with messages")
		},
	}
	server.ClientsMu.Lock()
	server.Clients[id] = client
	server.ClientsMu.Unlock()
	return client
}

type Payload struct {
	Success bool   `json:"success"`
	Origin  string `json:"origin,omitempty"`
	Event   string `json:"event"`
	Type    string `json:"type"`
	// Data    string `json:"data,omitempty"`
}

type BinaryPayload struct {
	Success bool   `json:"success"`
	Origin  string `json:"origin,omitempty"`
	Event   string `json:"event"`
	// Data    []byte `json:"data,omitempty"`
}
