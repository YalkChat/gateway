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

func createAdmin(db *gorm.DB) (*chat.User, error) {
	// ! Hash for default admin's "admin" password in BCrypt, it will not be this and
	// ! not be set this way.
	adminAccountPwd := "$2a$14$QuxLu/0REKoTuZGcwZwX2eLsNKFrook.QMh/Esd8d4FocaE2sKHca"
	adminAccount := &chat.Account{Email: "admin@example.com", Username: "admin", Password: adminAccountPwd, Verified: true}
	err := adminAccount.Create(db)
	if err != nil {
		logger.Err("CORE", fmt.Sprintf("FATAL - Can't create admin credentials, error: %v", err))
		return nil, err
	}
	logger.Info("CORE", fmt.Sprintf("Created admin credentials: %v", adminAccount.Username))

	adminUser := &chat.User{Account: adminAccount, DisplayedName: "Admin", AvatarUrl: "/default.png"}
	err = adminUser.Create(db)
	if err != nil {
		logger.Err("CORE", fmt.Sprintf("FATAL - Can't create admin profile, error: %v", err.Error()))
		return nil, err
	}
	logger.Info("CORE", fmt.Sprintf("Created admin user: %v", adminUser.DisplayedName))
	return adminUser, nil
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

func createBotUser(db *gorm.DB) error {
	botAccountPwd := "none"
	botAccount := &chat.Account{Email: "invalid@example.com", Username: "bot", Password: botAccountPwd, Verified: false}
	err := botAccount.Create(db)
	if err != nil {
		logger.Err("CORE", fmt.Sprintf("FATAL - Can't create bot credentials, error: %v", err))
		return nil
	}
	logger.Info("CORE", fmt.Sprintf("Created bot credentials: %v", botAccount.Username))
	serverBot := &chat.User{DisplayedName: "Bot", AvatarUrl: "/bot.png", Account: botAccount}
	err = serverBot.Create(db)
	if err != nil {
		return err
	}
	return nil
}
