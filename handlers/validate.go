package handlers

import (
	"log"
	"net/http"
	"yalk/cattp"
	"yalk/chat/server"
)

var ValidateHandle = cattp.HandlerFunc[*server.Server](func(w http.ResponseWriter, r *http.Request, context *server.Server) {
	defer r.Body.Close()
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := context.SessionsManager.Validate(context.Db, r, "YLK")
	if err != nil {
		// TODO: Extend session upon device validation
		log.Println("Invalid session")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// session.SetClientCookie(w)
	// TODO: Post response for WebSock?
	w.Header().Add("Content-Type", "application/json")
	// logger.Info("SESS", fmt.Sprintf("[%d][ID %v] Validated Session", http.StatusOK, session.AccountID))
	w.WriteHeader(http.StatusOK)
})
