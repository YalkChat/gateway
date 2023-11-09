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
	"fmt"
	"gateway/config"
	"gateway/http_server"

	"github.com/YalkChat/database"

	"github.com/joho/godotenv"
)

func init() {
	// * Clear console is kept only for debug reasons
	// * Won't be in the release version
	// log.Print("\033[H\033[2J") // Clear console
	var version string = "alpha-0.2"
	fmt.Printf("version: %s\n", version) // make it os.env
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("failed to load .env: %v", err)
		panic(err)
	}
	fmt.Println("loaded env variables")

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("failed to load ws config: %v", err)
		return
	}
	fmt.Println("WebSocket Config loaded")

	dbConfig, err := database.LoadConfig()
	if err != nil {
		fmt.Printf("failed to load db config: %v", err)
		return
	}
	fmt.Println("Database Config loaded")

	dbConn, err := database.InitializeDb(dbConfig)
	if err != nil {
		fmt.Printf("failed to inizialize db: %v", err)
		return
	}
	fmt.Println("DB connection initialized")

	httpConfig, err := http_server.LoadConfig()
	if err != nil {
		fmt.Printf("failed to load http config: %v", err)
		return
	}
	fmt.Println("HTTP config loaded")

	fmt.Println("running server")
	runServer(httpConfig, config, dbConn)
}
