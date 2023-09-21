package handlers

import (
	"fmt"
	"net/http"
	"time"
	"yalk/app"
	"yalk/chat/client"
	"yalk/chat/server"
	"yalk/sessions"

	"github.com/AleRosmo/cattp"
	"nhooyr.io/websocket"
)

func validateSession(r *http.Request, sessionsManager sessions.SessionManager) (*sessions.Session, error) {
	session, err := sessionsManager.Validate(r)
	if err != nil {
		return nil, err
	}
	return session, nil
}

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
	session, err := sessionsManager.Validate(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	conn, err := upgradeToWebSocket(w, r)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user, err := server.GetUserByID(session.UserID)
	if err != nil {
		handleError(w, err)
		return
	}

	userID := user.ID
	client, err := registerNewClient(conn, userID, server)
	if err != nil {
		handleError(w, err)
		return
	}
	// TODO: Consider if the initial payload should be done by the server instead
	err = sendInitialPayload(server, client.ID())
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Printf("registered new client: %d", client.ID())

})
