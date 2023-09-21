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

// TODO: Do these functions that only returns even make sense?

func validateSession(r *http.Request, sessionsManager sessions.SessionManager) (*sessions.Session, error) {
	return sessionsManager.Validate(r)
}

func upgradeToWebSocket(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return websocket.Accept(w, r, nil)
}

func registerNewClient(client client.Client, srv server.Server) error {
	return srv.RegisterClient(client)
}

// TODO: Placeholder, finish implementation
func handleError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)

}

func sendInitialPayload(srv server.Server, clientID uint) error {
	return nil
}

// TODO: Handle Error must be finished
var ConnectionHandler = cattp.HandlerFunc[app.HandlerContext](func(w http.ResponseWriter, r *http.Request, ctx app.HandlerContext) {
	server := ctx.ChatServer()
	sessionsManager := ctx.SessionsManager()

	session, err := validateSession(r, sessionsManager)
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

	// TODO: More correct to use session.UserID or user.ID()?
	userID := user.ID

	client := client.NewClient(userID, conn, time.Second*5) // TODO: time placeholder
	if err != nil {
		handleError(w, err)
	}

	err = registerNewClient(client, server)
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
