package http_server

import (
	"gateway/app"
	"gateway/handlers"
	"log"

	"github.com/AleRosmo/cattp"
)

func StartHttpServer(config *Config, context app.HandlerContext) error {
	router := cattp.New(context)

	router.HandleFunc("/ws", handlers.ConnectionHandler)

	// TODO: Temporarily disabled
	// router.HandleFunc("/auth", handlers.ValidateHandle)
	// router.HandleFunc("/auth/validate", handlers.ValidateHandle)
	// router.HandleFunc("/auth/signin", handlers.SigninHandle)
	// router.HandleFunc("/auth/signout", handlers.SignoutHandle)

	// router.HandleFunc("/auth/signup", signupHandle)

	netConf := cattp.Config{
		Host: config.Host,
		Port: config.Port,
		URL:  config.Url}

	err := router.Listen(&netConf)
	if err != nil {
		return err
	}

	log.Println("HTTP Server succesfully started") // TODO: Move back in main func
	return nil
}
