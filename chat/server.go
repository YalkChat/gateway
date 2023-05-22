package chat

import (
	"sync"
	"time"
	"yalk/chat/clients"
	"yalk/sessions"

	"golang.org/x/time/rate"
	"gorm.io/gorm"
	"nhooyr.io/websocket"
)

type ChatServer interface {
	RegisterClient(*websocket.Conn, string)
	SendMessage(*RawEvent)
	SendMessageToAll(*RawEvent)
	Sender(*clients.Client, *EventContext)
	Receiver(*EventContext)
	HandlePayload([]byte)
}

// TODO: db
func NewServer(bufferLenght uint, db *gorm.DB, sessionsManager *sessions.Manager) *Server {
	// sessionLen := time.Hour * 720
	// sessionsManager := sessions.New(sessionLen)

	sendLimiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 8)
	clientsMap := make(map[uint]*clients.Client)
	messageChannels := MakeEventChannels()

	chatServer := &Server{
		SendLimiter:          sendLimiter,
		Clients:              clientsMap,
		ClientsMessageBuffer: bufferLenght,
		Channels:             messageChannels,
		Db:                   db,
		SessionsManager:      sessionsManager,
	}

	return chatServer
}

type Server struct {
	SendLimiter          *rate.Limiter
	Clients              map[uint]*clients.Client
	ClientsMu            sync.Mutex
	ClientsMessageBuffer uint
	Channels             *EventChannels
	Db                   *gorm.DB
	SessionsManager      *sessions.Manager
}

func (server *Server) RegisterClient(conn *websocket.Conn, id uint) *clients.Client {
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

type BinaryPayload struct {
	Success bool   `json:"success"`
	Origin  string `json:"origin,omitempty"`
	Event   string `json:"event"`
	// Data    []byte `json:"data,omitempty"`
}

type ServerSettings struct {
	gorm.Model
	IsInitialized bool `json:"is_initialized"`
}

func (s *ServerSettings) Create(db *gorm.DB) error {
	return db.Create(s).Error
}

func (s *ServerSettings) Update(db *gorm.DB) error {
	return db.Save(s).Error
}
