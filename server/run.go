package server

import (
	"fmt"
	"gateway/app"
	"gateway/config"
	"gateway/encryption"
	"sync"
	"time"

	"gateway/sessions"

	"github.com/AleRosmo/engine/server"

	"github.com/AleRosmo/engine/initialize"
	"github.com/AleRosmo/engine/serialization/strategies"

	"github.com/AleRosmo/engine/database"

	"gorm.io/gorm"
)

func RunServer(config *config.Config, conn *gorm.DB) {
	var wg sync.WaitGroup
	var sessionLenght = time.Hour * 720
	var cookieName = "YLK"

	sessionsDatabase := sessions.NewDatabase(conn)
	encryptionService := &encryption.BcryptService{}
	sessionsManager := sessions.NewSessionManager(sessionsDatabase, encryptionService, sessionLenght, cookieName)

	// chatServer := server.NewServer(16, conn, sessionsManager)

	// Initialize the serialization strategy
	serializationStrategy := &strategies.JSONStrategy{} // or &serialization.ProtobufStrategy{}

	// Initialize db operations
	dbOperations := database.NewDatabase(conn)

	// TODO: the name suggests it might not be in the rightr package
	if err := initialize.InitializeApp(dbOperations); err != nil {
		fmt.Printf("failed to inizialize app: %v", err)
		return
	}
	newChatServer := server.NewServer(dbOperations, serializationStrategy)
	// TODO: Add a factory function for app.HandlerContext
	// TODO: It's not ok to call the method like this
	// TODO: CHANGE REALLY IT'S HORRIBLE
	handlerContext := app.NewHandlerContext(newChatServer, sessionsManager, config)

	wg.Add(1)
	go func() {
		defer wg.Done()
		StartHttpServer(config, handlerContext)
	}()

	fmt.Println("server started")
	wg.Wait()
}
