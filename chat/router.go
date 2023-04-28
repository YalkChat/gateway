package chat

import (
	"encoding/json"
	"fmt"
	"log"
)

func (server *Server) SendMessage(senderId string, jsonPayload []byte, echoToSender bool) {
	for userId, client := range server.Clients {
		if userId != senderId || echoToSender {
			client.Msgs <- jsonPayload
		}
	}
}

func (server *Server) Router() {
	for {
		select {
		case payload := <-server.Channels.Conn:
			fmt.Println("Router: Conn received")
			jsonPayload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			server.SendMessage(payload.Origin, jsonPayload, false)

		case payload := <-server.Channels.Disconn:
			jsonPayload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			server.SendMessage(payload.Origin, jsonPayload, false)

		case payload := <-server.Channels.Msg:
			jsonPayload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}
			server.SendMessage(payload.Origin, jsonPayload, true)

		case _p := <-server.Channels.Dm:
			fmt.Println("Router: Dm received")
			dest := _p["users"].([]string)
			payload := _p["payload"].(Payload)
			_payload, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Marshaling err")
			}

			for _, id := range dest {
				wsClient := server.Clients[id]
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
			for _, wsClient := range server.Clients {
				wsClient.Msgs <- payload
			}
		}
	}
}
