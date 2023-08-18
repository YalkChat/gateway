package main

import (
	"fmt"
	"sync"
	"time"
	"yalk/cattp"
	"yalk/chat"
	"yalk/sessions"

	"gorm.io/gorm"
)

func runServer(config *Config, db *gorm.DB) {
	var wg sync.WaitGroup

	netConf := cattp.Config{
		Host: config.HttpHost,
		Port: config.HttpPort,
		URL:  config.HttpUrl}

	sessionLenght := time.Hour * 720
	sessionsManager := sessions.New(db, sessionLenght)

	chatServer := chat.NewServer(16, db, sessionsManager)

	wg.Add(1)
	go func() {
		defer wg.Done()
		chatServer.Router()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		startHttpServer(netConf, chatServer)
	}()

	fmt.Println("server started")
	wg.Wait()
}
