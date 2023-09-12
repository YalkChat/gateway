package app

import (
	"yalk/chat/server"
	"yalk/sessions"
)

type handlerContextImpl struct {
	server          server.Server
	sessionsManager sessions.SessionManager
}

func NewHandlerContext(server server.Server, sessionsManager sessions.SessionManager) HandlerContext {
	return &handlerContextImpl{server: server, sessionsManager: sessionsManager}

}

func (hc *handlerContextImpl) ChatServer() server.Server {
	return hc.server
}

func (hc *handlerContextImpl) SessionsManager() sessions.SessionManager {
	return hc.sessionsManager
}
