package chat

import (
	"database/sql"
	"sync"
	"time"

	"golang.org/x/time/rate"
	"nhooyr.io/websocket"
)

// TODO: db
func NewServer(bufferLenght int) *Server {

	sendLimiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 8)
	clientsMap := make(map[string]*Client)
	messageChannels := makeMessageChannels()

	chatServer := &Server{
		SendLimiter:          sendLimiter,
		Clients:              clientsMap,
		ClientsMessageBuffer: bufferLenght,
		Channels:             messageChannels,
		// db: db,
	}

	return chatServer
}

type ChatServer interface {
	RegisterClient()
}

type Server struct {
	SendLimiter          *rate.Limiter
	Clients              map[string]*Client
	ClientsMu            sync.Mutex
	ClientsMessageBuffer int
	Channels             MessageChannels
	db                   *sql.DB
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
