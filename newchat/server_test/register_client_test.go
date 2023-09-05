package server_test

// TODO: Doesn't work anymore, readapt
// TODO: Commented out but requires urgent review

// func TestRegisterClient(t *testing.T) {
// 	// Initialize a new server
// 	// TODO: Should the Database contain an actual connection?
// 	srv := server.NewServer(&database.DatabaseImpl{}, &sessions.Sess)

// 	// Create a mock client
// 	client := client.NewClient("123", &websocket.Conn{})

// 	// Register the client
// 	srv.RegisterClient(client)

// 	client, err := srv.GetClientByID(client.ID())
// 	if err != nil {
// 		t.Errorf("Client with ID %s was not registered: %v", client.ID(), err)
// 	}
// }
