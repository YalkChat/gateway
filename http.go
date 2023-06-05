package main

import (
	"log"
	"yalk/cattp"
	"yalk/chat"
	"yalk/handlers"
)

func startHttpServer(conf cattp.Config, chatServer *chat.Server) error {
	router := cattp.New(chatServer)

	router.HandleFunc("/ws", handlers.ConnectHandle)

	router.HandleFunc("/auth", validateHandle)
	router.HandleFunc("/auth/validate", validateHandle)
	router.HandleFunc("/auth/signin", signinHandle)
	router.HandleFunc("/auth/signout", signoutHandle)
	router.HandleFunc("/auth/signup", signupHandle)

	err := router.Listen(&conf)
	if err != nil {
		return err
	}

	log.Println("HTTP Server succesfully started") // TODO: Move back in main func
	return nil
}
