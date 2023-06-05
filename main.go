package main

// ** Server events and meaning ** //

// ** - 'user_login' -- User connecting to server
// ** - 'user_disconn' -- User disconnecting from server
// ** - 'user_new' -- New user account
// ** - 'user_delete' -- User account deleted
// ** - 'user_update' -- User info update

// ** - 'chat_create' -- New Chat created
// ** - 'chat_delete' -- Chat deleted
// ** - 'chat_message' -- Chat message received
// ** - 'chat_join' -- Chat joined by another user
// ** - 'chat_invite' -- Chat invite received by another user
// ** - 'chat_leave' -- Chat left by another user

import (
	"time"
	"yalk/cattp"
	"yalk/chat"
	"yalk/logger"
	"yalk/sessions"

	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func init() {
	log.Print("\033[H\033[2J") // Clear console
	var version string = "pre-alpha"
	logger.Info("CORE", "Booting..")
	logger.Info("CORE", fmt.Sprintf("Chat Server version: %s", version)) // make it os.env
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

}

func main() {
	var wg sync.WaitGroup

	// ? Separate function
	dbConfig := &chat.PgConf{
		IP:       os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := chat.OpenDbConnection(dbConfig)
	if err != nil {
		logger.Err("CORE", fmt.Sprintf("Error opening db connection: %v\n", err))
		return
	}

	if err := chat.CreateDbTables(db); err != nil {
		logger.Warn("DB", fmt.Sprintf("Failed to AutoMigrate DB tables: %v", err))
	}

	if isInitialized := checkIsInitialized(db); !isInitialized {

		if err := createBotUser(db); err != nil {
			logger.Err("CORE", fmt.Sprintf("FATAL - Can't create initial DB user, error: %v", err.Error()))
		}
		logger.Info("CORE", "Created Bot user")

		adminUser, err := createAdmin(db)
		if err != nil {
			return
		}
		logger.Info("CORE", "Admin creation succesful")

		chatType := &chat.ChatType{Type: "channel"}
		tx := db.Create(chatType)
		if tx.Error != nil {
			logger.Err("CORE", fmt.Sprintf("FATAL - Can't create chat type, error: %v", tx.Error))
		}
		logger.Info("CORE", "Channel type creation succesful")

		mainChat := &chat.Chat{Name: "Main", ChatType: chatType, CreatedBy: adminUser, Users: []*chat.User{adminUser}}

		tx = db.Create(mainChat)
		if tx.Error != nil {
			logger.Err("CORE", fmt.Sprintf("FATAL - Can't create main chat, error: %v", tx.Error))
		}
		logger.Info("CORE", "Main chat creation succesful")

		serverSettings := &chat.ServerSettings{IsInitialized: true}
		serverSettings.Create(db)
		logger.Info("CORE", "Created initial DB settings")
	}

	netConf := cattp.Config{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT_PLAIN"),
		URL:  os.Getenv("HTTP_URL"),
	}
	sessionLenght := time.Hour * 720
	sessionsManager := sessions.New(db, sessionLenght)

	chatServer := chat.NewServer(16, db, sessionsManager)

	wg.Add(1)
	go chatServer.Router()

	wg.Add(1)
	go startHttpServer(netConf, chatServer)

	wg.Wait()
}
