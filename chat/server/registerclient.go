package server

import (
	"yalk/chat/chatmodels"

	"nhooyr.io/websocket"
)

func (server *Server) RegisterClient(conn *websocket.Conn, id uint) *chatmodels.Client {

	// if client, ok := server.Clients[id]; ok {
	// 	logger.Info("SRV", fmt.Sprintf("Client already registerd: %d", id))
	// 	return client
	// }

	messageChan := make(chan []byte, server.ClientsMessageBuffer)

	client := &chatmodels.Client{
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
