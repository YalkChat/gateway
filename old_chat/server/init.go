package server

import (
	"sync"
	"time"
	"yalk/chat/models"

	"yalk/sessions"

	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

// type ChatServer interface {
// 	RegisterClient(*websocket.Conn, string)
// 	SendMessage(*models.RawEvent)
// 	SendMessageToAll(*models.RawEvent)
// 	Sender(*models.Client, *models.EventContext)
// 	Receiver(*models.EventContext)
// 	HandlePayload([]byte)
// }

// TODO: db
func NewServer(bufferLenght uint, db *gorm.DB, sessionsManager sessions.SessionManager) *Server {
	// sessionLen := time.Hour * 720
	// sessionsManager := sessions.New(sessionLen)

	sendLimiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 8)
	clientsMap := make(map[uint]*models.Client)

	chatServer := &Server{
		SendLimiter:          sendLimiter,
		Clients:              clientsMap,
		ClientsMessageBuffer: bufferLenght,
		Db:                   db,
		SessionsManager:      sessionsManager,
	}

	return chatServer
}

type Server struct {
	SendLimiter          *rate.Limiter
	Clients              map[uint]*models.Client
	ClientsMu            sync.Mutex
	ClientsMessageBuffer uint
	Db                   *gorm.DB
	SessionsManager      sessions.SessionManager
}

type BinaryPayload struct {
	Success bool   `json:"success"`
	Origin  string `json:"origin,omitempty"`
	Event   string `json:"event"`
	// Data    []byte `json:"data,omitempty"`
}
