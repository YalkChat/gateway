package appcontext

import (
	"yalk/chat/server"
	"yalk/sessions"
)

type handlerContextImpl struct {
	server          server.Server
	sessionsManager sessions.SessionManager
}

func (hc *handlerContextImpl) Server() server.Server {
	return hc.server
}

func (hc *handlerContextImpl) SessionsManager() sessions.SessionManager {
	return hc.sessionsManager
}
