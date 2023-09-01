package server_test

import (
	"testing"
	"yalk/database"
	"yalk/newchat/client"
	"yalk/newchat/server"

	"nhooyr.io/websocket"
)

func TestRegisterClient(t *testing.T) {
	// Initialize a new server
	srv := server.NewServer(&database.DatabaseImpl{})

	// Create a mock client
	client := client.NewClient("123", &websocket.Conn{})

	// Register the client
	srv.RegisterClient(client)

	client, err := srv.GetClientByID(client.ID())
	if err != nil {
		t.Errorf("Client with ID %s was not registered: %v", client.ID(), err)
	}
}
