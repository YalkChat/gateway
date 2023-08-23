package server

import (
	"errors"
	"yalk/chat/models"
)

func (server *Server) UnregisterClient(c *models.Client) error {
	if server.Clients[c.ID] == nil {
		return errors.New("no_client")
	}
	server.ClientsMu.Lock()
	delete(server.Clients, c.ID)
	server.ClientsMu.Unlock()
	return nil
}
