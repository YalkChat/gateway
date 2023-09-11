package handlers

import (
	"fmt"
	"net/http"
	"yalk/app"
	"yalk/chat/client"
	"yalk/chat/server"

	"github.com/AleRosmo/cattp"
	"nhooyr.io/websocket"
)

// TODO: Removed as simplified in sessionsManager.Validate() below
// func validateSession(r *http.Request) (*sessions.Session, error) {

// }

func upgradeToWebSocket(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return nil, nil // TODO: First is a placeholder for conn, change
}

func registerNewClient(conn *websocket.Conn, userID string, srv *server.Server) (*client.Client, error) {
	return nil, nil // TODO: Placeholder, implement
}

func handleError(w http.ResponseWriter, err error) {
	return
}

func sendInitialPayload(srv *server.Server, clientID string) error {
	return nil
}

var ConnectionHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	server := ctx.ChatServer()
	sessionsManager := ctx.SessionManager()

	// TODO: remove _ placeholder below
	_, err := sessionsManager.Validate(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	conn, err := upgradeToWebSocket(w, r)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// TODO: Find ID here
	// TODO: 1 is placeholder
	client := client.NewClient(1, conn)

	err = server.RegisterClient(client)
	if err != nil {
		// TODO: placeholder
		return
	}
	fmt.Printf("registered new client: %d", client.ID())

})
