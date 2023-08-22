package server

import (
	"errors"
	"yalk/chat/clients"
)

func (server *Server) UnregisterClient(c *clients.Client) error {
	if server.Clients[c.ID] == nil {
		return errors.New("no_client")
	}
	server.ClientsMu.Lock()
	delete(server.Clients, c.ID)
	server.ClientsMu.Unlock()
	return nil
}
