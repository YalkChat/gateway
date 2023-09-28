package handlers

import (
	"yalk/app"
	"yalk/chat/server"
	"yalk/config"
	"yalk/sessions"
)

func getContextComponents(ctx app.HandlerContext) (server.Server, sessions.SessionManager, *config.Config) {
	return ctx.ChatServer(), ctx.SessionsManager(), ctx.Config()
}
