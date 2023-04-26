package chat

import (
	"encoding/json"
	"fmt"
	"log"
)

func (server *Server) Router() {
	for {
		select {
		case payload := <-server.Channels.Conn:
			fmt.Println("Router: Conn received")
			_payload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			for i, wsClient := range server.Clients {
				if i != payload.Origin {
					wsClient.Msgs <- _payload
				}
			}

		case payload := <-server.Channels.Disconn:
			_payload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			for i, wsClient := range server.Clients {
				if i != payload.Origin {
					wsClient.Msgs <- _payload
				}
			}

		case payload := <-server.Channels.Msg:
			_payload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			// for _, client_chan := range server.webserver.Clients {
			// 	client_chan <- _payload
			// }
			for _, wsClient := range server.Clients {
				wsClient.Msgs <- _payload
			}

		case _p := <-server.Channels.Dm:
			fmt.Println("Router: Dm received")
			dest := _p["users"].([]string)
			payload := _p["payload"].(Payload)
			_payload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}

			for _, id := range dest {
				// sseClient := server.webserver.SSEClients[id]
				wsClient := server.Clients[id]

				// if sseClient != nil {
				// 	sseClient <- _payload
				// }
				if wsClient != nil {
					wsClient.Msgs <- _payload
				}
			}

		case _payload := <-server.Channels.Notify:
			fmt.Println("Router: Notify received")
			payload, err := json.Marshal(_payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			// for _, client_chan := range server.SSEClients {
			// 	client_chan <- payload
			// }
			for _, wsClient := range server.Clients {
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
