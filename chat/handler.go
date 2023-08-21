package chat

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// * Handle incoming user payload and process it eventually
// * forwarding in the correct routine channels for other users to receive.
func (server *Server) HandleIncomingEvent(clientID uint, rawEvent *RawEvent) error {
	log.Printf("Handling incoming event for ID %d", clientID)
	switch rawEvent.Type {
	case "message":
		newMessage, err := handleMessage(rawEvent, server.Db)
		if err != nil {
			return err
		}
		server.Channels.Messages <- newMessage

	case "user_login":
		server.Channels.Notify <- rawEvent

	// ? Non forwarded events - server only
	case "account":
		newAccount, err := handleAccount(rawEvent, server.Db)
		if err != nil {
			return err
		}
		// // TODO: To move to initial profile setup
		// newUser, err := handleUserCreate(rawEvent, server.Db, newAccount)
		// if err != nil {
		// 	return err
		// }
		serializedData, err := newAccount.Serialize()
		if err != nil {
			log.Printf("Error serializing")
		}

		rawEvent.Data = serializedData

		server.Channels.Accounts <- rawEvent

	case "user":
		updateUser, err := handleUser(rawEvent, server.Db)
		if err != nil {
			return err
		}
		serializedData, err := updateUser.Serialize()
		if err != nil {
			log.Printf("Error serializing")
		}

		rawEvent.Data = serializedData
		server.Channels.Users <- rawEvent

	case "initial":
		fmt.Println("ok intiial")

	default:
		log.Printf("Payload Handler received an invalid event type")
		return fmt.Errorf("invalid_request")
	}
	return nil
}

func handleMessage(rawEvent *RawEvent, db *gorm.DB) (*Message, error) {
	var user *User
	var message *Message
	var chat *Chat

	switch rawEvent.Action {
	case "send":
		user = &User{ID: rawEvent.UserID}
		if err := user.GetInfo(db); err != nil {
			log.Printf("Error getting user info ID: %d\n", rawEvent.UserID)
			return nil, err

		}

		message = &Message{UserID: rawEvent.UserID}
		if err := message.Deserialize(rawEvent.Data); err != nil {
			log.Printf("Error Deserializing Chat Message")
			return nil, err
		}

		// chat = &Chat{ID: message.ChatID}
		// if _, err := chat.GetInfo(db); err != nil {
		// 	log.Printf( fmt.Sprintf("Error getting message chat ID: %d\n", message.UserID))
		// 	return nil, err
		// }

		message.Chat = chat

		if err := message.SaveToDb(db); err != nil {
			log.Printf("Error saving to DB Chat Message")
			return nil, err
		}

	}
	return message, nil
}

func handleAccount(rawEvent *RawEvent, db *gorm.DB) (*Account, error) {
	account := &Account{}

	if err := account.Deserialize(rawEvent.Data); err != nil {
		log.Printf("Error Deserializing Account")
		return nil, err
	}

	if err := account.Create(db); err != nil {
		log.Printf("Error Creating Account: %d\n", account.ID)
		return nil, err
	}

	log.Printf("Account Created: %d\n", account.ID)
	return account, nil
}

// func handleUserCreate(rawEvent *RawEvent, db *gorm.DB, account *Account) (*User, error) {
// 	user := &User{Account: account}

// 	if err := user.Deserialize(rawEvent.Data); err != nil {
// 		log.Printf( "Error Deserializing User")
// 		return nil, err
// 	}

// 	if err := user.Create(db); err != nil {
// 		log.Printf( fmt.Sprintf("Error Creating User: %d\n", user.ID))
// 		return nil, err
// 	}

// 	log.Printf("User Created: %d\n", user.ID))
// 	return user, nil
// }

func handleUser(rawEvent *RawEvent, db *gorm.DB) (*User, error) {
	var newUser = &User{ID: rawEvent.UserID}
	// var status = &Status{}
	if err := newUser.GetInfo(db); err != nil {
		log.Printf("Error getting user info ID: %d\n", newUser.ID)
		return nil, err
	}
	switch rawEvent.Action {
	case "change_status":

		// TODO: Change to status event type
		statusPayload := &User{}
		if err := statusPayload.Deserialize(rawEvent.Data); err != nil {
			log.Printf("Error Deserializing User")
			return nil, err
		}

		newUser.Status = &Status{Name: statusPayload.StatusName}

		if err := newUser.SaveToDb(db); err != nil {
			log.Printf(fmt.Sprintf("Error saving to DB User: %d\n", newUser.ID))
			return nil, err
		}
	}
	return newUser, nil
}
