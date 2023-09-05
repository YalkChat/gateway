package server

import (
	"fmt"
	"sync"
	"time"
	"yalk/chat/server"
	"yalk/config"
	"yalk/database"
	newserver "yalk/newchat/server"
	"yalk/sessions"

	"gorm.io/gorm"
)

func RunServer(config *config.Config, conn *gorm.DB) {
	var wg sync.WaitGroup

	sessionDatabase := sessions.NewDatabase()

	sessionLenght := time.Hour * 720
	sessionsManager := sessions.NewSessionManager(sessionDatabase, sessionLenght)

	chatServer := server.NewServer(16, conn, sessionsManager)
	db := database.NewDatabase(conn)
	newChatServer := newserver.NewServer(db)
	fmt.Print(newChatServer) // TODO: remove
	wg.Add(1)
	go func() {
		defer wg.Done()
		StartHttpServer(config, chatServer)
	}()

	fmt.Println("server started")
	wg.Wait()
}
