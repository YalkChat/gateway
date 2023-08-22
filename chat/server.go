package chat

import (
	"errors"
	"sync"
	"time"
	"yalk/chat/chatmodels"
	"yalk/chat/clients"
	"yalk/sessions"

	"golang.org/x/time/rate"
	"gorm.io/gorm"
	"nhooyr.io/websocket"
)

type ChatServer interface {
	RegisterClient(*websocket.Conn, string)
	SendMessage(*chatmodels.RawEvent)
	SendMessageToAll(*chatmodels.RawEvent)
	Sender(*clients.Client, *chatmodels.EventContext)
	Receiver(*chatmodels.EventContext)
	HandlePayload([]byte)
}

// TODO: db
func NewServer(bufferLenght uint, db *gorm.DB, sessionsManager *sessions.Manager) *Server {
	// sessionLen := time.Hour * 720
	// sessionsManager := sessions.New(sessionLen)

	sendLimiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 8)
	clientsMap := make(map[uint]*clients.Client)
	messageChannels := chatmodels.MakeEventChannels() // TODO Move this to server package
	messageMap := make(map[uint]bool)

	chatServer := &Server{
		SendLimiter:          sendLimiter,
		Clients:              clientsMap,
		ClientsMessageBuffer: bufferLenght,
		Channels:             messageChannels,
		Db:                   db,
		SessionsManager:      sessionsManager,
		MessageMap:           messageMap,
	}

	return chatServer
}

type Server struct {
	SendLimiter          *rate.Limiter
	Clients              map[uint]*clients.Client
	ClientsMu            sync.Mutex
	ClientsMessageBuffer uint
	Db                   *gorm.DB
	SessionsManager      *sessions.Manager
	MessageMap           map[uint]bool
}

func (server *Server) RegisterClient(conn *websocket.Conn, id uint) *clients.Client {

	// if client, ok := server.Clients[id]; ok {
	// 	logger.Info("SRV", fmt.Sprintf("Client already registerd: %d", id))
	// 	return client
	// }

	messageChan := make(chan []byte, server.ClientsMessageBuffer)

	client := &clients.Client{
		ID:   id,
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

func (server *Server) UnregisterClient(c *clients.Client) error {
	if server.Clients[c.ID] == nil {
		return errors.New("no_client")
	}
	server.ClientsMu.Lock()
	delete(server.Clients, c.ID)
	server.ClientsMu.Unlock()
	return nil
}

type BinaryPayload struct {
	Success bool   `json:"success"`
	Origin  string `json:"origin,omitempty"`
	Event   string `json:"event"`
	// Data    []byte `json:"data,omitempty"`
}
