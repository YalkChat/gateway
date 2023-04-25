package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type server struct {
	// instance int
	// config   NetworkConfig
	// settings Settings
	channels  channels
	websocket *webSocketServer
	dbconn    *sql.DB
}

func (server *server) router() {
	for {
		select {
		case payload := <-server.channels.Conn:
			fmt.Println("Router: Conn received")
			_payload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			for i, wsClient := range server.websocket.Clients {
				if i != payload.Origin {
					wsClient.Msgs <- _payload
				}
			}

		case payload := <-server.channels.Disconn:
			_payload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			for i, wsClient := range server.websocket.Clients {
				if i != payload.Origin {
					wsClient.Msgs <- _payload
				}
			}

		case payload := <-server.channels.Msg:
			_payload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			// for _, client_chan := range server.webserver.Clients {
			// 	client_chan <- _payload
			// }
			for _, wsClient := range server.websocket.Clients {
				wsClient.Msgs <- _payload
			}

		case _p := <-server.channels.Dm:
			fmt.Println("Router: Dm received")
			dest := _p["users"].([]string)
			payload := _p["payload"].(payload)
			_payload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}

			for _, id := range dest {
				// sseClient := server.webserver.SSEClients[id]
				wsClient := server.websocket.Clients[id]

				// if sseClient != nil {
				// 	sseClient <- _payload
				// }
				if wsClient != nil {
					wsClient.Msgs <- _payload
				}
			}

		case _payload := <-server.channels.Notify:
			fmt.Println("Router: Notify received")
			payload, err := json.Marshal(_payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			// for _, client_chan := range server.SSEClients {
			// 	client_chan <- payload
			// }
			for _, wsClient := range server.websocket.Clients {
				wsClient.Msgs <- payload
			}

			// case payload := <-server.channels.Cmd:
			// 	// TODO: Check user admin status
			// fields := strings.Fields(payload.Data["message"].(map[string]string))
			// command := fields[0]
			// srvMessage := fields[1]
			// //? CHANGE PLACE?
			// switch command {
			// case "/test":
			// 	payload.Event = "server_message"
			// 	payload.Message.Text = srvMessage
			// 	server.channels.Notify <- payload
		}
	}
}
