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
	"yalk/initialize"

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

	config, err := loadConfig()
	if err != nil {
		fmt.Printf("failed to load config: %v", err)
		return
	}
	fmt.Println("Config loaded")

	db, err := initializeDb(config)
	if err != nil {
		fmt.Printf("failed to inizialize db: %v", err)
		return
	}
	fmt.Println("DB connection initialized")

	if err := initialize.InitializeApp(db); err != nil {
		fmt.Printf("failed to inizialize app: %v", err)
		return
	}
	fmt.Println("app initialized")

	runServer(config, db)
}
