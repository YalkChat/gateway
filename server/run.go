package server

import (
	"fmt"
	"sync"
	"time"
	"yalk/config"
	"yalk/serialization"

	"yalk/chat/database"
	"yalk/chat/initialize"
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

	// Initialize the serialization strategy
	serializationStrategy := &serialization.JSONStrategy{} // or &serialization.ProtobufStrategy{}

	db := database.NewDatabase(conn)

	if err := initialize.InitializeApp(db); err != nil {
		fmt.Printf("failed to inizialize app: %v", err)
		return
	}
	newChatServer := server.NewServer(db, sessionsManager, serializationStrategy)
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
