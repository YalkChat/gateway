package app

import (
	"yalk/chat/server"
	"yalk/config"
	"yalk/sessions"
)

type handlerContextImpl struct {
	server          server.Server
	sessionsManager sessions.SessionManager
	config          *config.Config //TODO: Interface here instead makes sense?
}

func NewHandlerContext(server server.Server, sessionsManager sessions.SessionManager, config *config.Config) HandlerContext {
	return &handlerContextImpl{
		server:          server,
		sessionsManager: sessionsManager,
		config:          config,
	}
}

func (hc *handlerContextImpl) ChatServer() server.Server {
	return hc.server
}

func (hc *handlerContextImpl) SessionsManager() sessions.SessionManager {
	return hc.sessionsManager
}

func (hc *handlerContextImpl) Config() *config.Config {
	return hc.config
}
