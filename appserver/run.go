package appserver

import (
	"fmt"
	"sync"
	"time"
	"yalk/chat/server"
	"yalk/config"
	newserver "yalk/newchat/server"
	"yalk/sessions"

	"gorm.io/gorm"
)

func RunServer(config *config.Config, db *gorm.DB) {
	var wg sync.WaitGroup

	sessionLenght := time.Hour * 720
	sessionsManager := sessions.New(db, sessionLenght)

	chatServer := server.NewServer(16, db, sessionsManager)
	newChatServer := newserver.NewServer(db)
	fmt.Print(newChatServer)
	wg.Add(1)
	go func() {
		defer wg.Done()
		StartHttpServer(config, chatServer)
	}()

	fmt.Println("server started")
	wg.Wait()
}
