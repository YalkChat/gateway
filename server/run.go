package server

import (
	"fmt"
	"sync"
	"time"
	"yalk/config"

	"yalk/chat/database"
	"yalk/chat/server"
	"yalk/sessions"

	"gorm.io/gorm"
)

func RunServer(config *config.Config, conn *gorm.DB) {
	var wg sync.WaitGroup
	var sessionLenght = time.Hour * 720
	var cookieName = "YLK"

	sessionsDatabase := sessions.NewDatabase(conn)
	sessionsManager := sessions.NewSessionManager(sessionsDatabase, sessionLenght, cookieName)

	// chatServer := server.NewServer(16, conn, sessionsManager)
	db := database.NewDatabase(conn)
	newChatServer := server.NewServer(db, sessionsManager)
	fmt.Print(newChatServer) // TODO: remove
	// TODO: enable again when modifying StartHttpServer()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	StartHttpServer(config, newChatServer)
	// }()

	fmt.Println("server started")
	wg.Wait()
}
