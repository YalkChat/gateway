package handlers

import (
	"fmt"
	"net/http"
	"time"
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
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil // TODO: First is a placeholder for conn, change
}

func registerNewClient(conn *websocket.Conn, userID uint, srv server.Server) (client.Client, error) {
	newClient := client.NewClient(userID, conn, time.Second*5) // TODO: time placeholder
	err := srv.RegisterClient(newClient)
	if err != nil {
		return nil, err
	}
	return newClient, nil
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)

}

func sendInitialPayload(srv server.Server, clientID uint) error {
	return nil
}

var ConnectionHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	server := ctx.ChatServer()
	sessionsManager := ctx.SessionsManager()

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
	var userID uint = 1 // This should be replaced with actual logic to find the userID

	client, err := registerNewClient(conn, userID, server)
	if err != nil {
		handleError(w, err)
		return
	}
	err = sendInitialPayload(server, client.ID())
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Printf("registered new client: %d", client.ID())

})
