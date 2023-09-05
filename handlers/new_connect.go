package handlers

import (
	"net/http"
	"yalk/chat/server"
	"yalk/newchat/client"
	"yalk/sessions"

	"github.com/AleRosmo/cattp"
	"nhooyr.io/websocket"
)

func validateSession(r *http.Request) (*sessions.Session, error) {
	return nil, nil
}

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

var ConnectionHandler = cattp.HandlerFunc[*server.Server](func(w http.ResponseWriter, r *http.Request, server *server.Server) {
	sessionsManager := server.SessionsManager
	cookieName := "YLK"

	session, err := sessionsManager.Validate(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return err
	}

	conn, err := upgradeToWebSocket(w, r)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	clientID := srv.RegisterClient(conn, session.AccountID)
	return nil
})
