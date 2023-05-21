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
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"nhooyr.io/websocket"
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
		serverBot := &chat.User{Username: "bot", Email: "admin@example.com", DisplayedName: "Bot", AvatarUrl: "/bot.png"}
		_, err := serverBot.Create(db)
		if err != nil {
			logger.Err("CORE", fmt.Sprintf("FATAL - Can't create initial DB user, error: %v", err.Error()))
			return
		}
		logger.Info("CORE", fmt.Sprintf("Created initial DB user: %v", serverBot.Username))

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

func upgradeHttpRequest(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	var defaultOptions = &websocket.AcceptOptions{CompressionMode: websocket.CompressionNoContextTakeover, InsecureSkipVerify: true}
	var defaultSize int64 = 2097152 // 2Mb in bytes

	conn, err := websocket.Accept(w, r, defaultOptions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		r.Body.Close()
		return nil, err
	}

	conn.SetReadLimit(defaultSize)
	return conn, nil
}

func checkIsInitialized(db *gorm.DB) bool {
	var serverSettings *chat.ServerSettings
	tx := db.Select("is_initialized").First(&serverSettings, "is_initialized = true")
	return tx.Error == nil
}
