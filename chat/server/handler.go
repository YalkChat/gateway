package server

import (
	"encoding/json"
	"fmt"
	"log"
	"yalk/chat/handlers"
	"yalk/chat/models"
)

// * Handle incoming user payload and process it eventually
// * forwarding in the correct routine channels for other users to receive.
func (server *Server) HandleIncomingEvent(clientID uint, rawEvent *models.RawEvent) error {
	log.Printf("Handling incoming event for ID %d", clientID)
	switch rawEvent.Type {
	case "message":
		// TODO: Change name of function, or refactor function
		newMessage, err := handlers.HandleMessage(rawEvent, server.Db)
		if err != nil {
			return err
		}
		serializedData, err := newMessage.Serialize()
		if err != nil {
			log.Printf("Error serializing")
		}

		newRawEvent := models.RawEvent{UserID: newMessage.UserID, Type: newMessage.Type(), Data: serializedData}

		jsonEvent, err := json.Marshal(newRawEvent)
		if err != nil {
			log.Printf("Error serializing RawEvent")
		}
		server.SendToChat(newMessage, jsonEvent)

	// case "user_login":
	// 	server.Channels.Notify <- rawEvent

	// // ? Non forwarded events - server only
	// case "account":
	// 	newAccount, err := handlers.HandleAccount(rawEvent, server.Db)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// // TODO: To move to initial profile setup
	// 	// newUser, err := handleUserCreate(rawEvent, server.Db, newAccount)
	// 	// if err != nil {
	// 	// 	return err
	// 	// }
	// 	serializedData, err := newAccount.Serialize()
	// 	if err != nil {
	// 		log.Printf("Error serializing")
	// 	}

	// 	rawEvent.Data = serializedData

	// 	server.Channels.Accounts <- rawEvent

	// case "user":
	// 	updateUser, err := handlers.HandleUser(rawEvent, server.Db)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	serializedData, err := updateUser.Serialize()
	// 	if err != nil {
	// 		log.Printf("Error serializing")
	// 	}

	// 	rawEvent.Data = serializedData
	// 	server.Channels.Users <- rawEvent

	// TODO: The goddamn password is sent with the initial payload
	case "initial":
		fmt.Println("ok intiial")

	default:
		log.Printf("Payload Handler received an invalid event type")
		return fmt.Errorf("invalid_request")
	}
	return nil
}
