package chat

import (
	"fmt"

	"yalk/logger"

	"gorm.io/gorm"
)

// * Handle incoming user payload and process it eventually
// * forwarding in the correct routine channels for other users to receive.
func (server *Server) HandleIncomingEvent(clientID uint, jsonEventMessage []byte) error {

	rawEvent := &RawEvent{UserID: clientID}

	if err := rawEvent.Deserialize(jsonEventMessage); err != nil {
		logger.Err("HNDL", "Error unmarshaling RawEvent")
		return err
	}

	switch rawEvent.Type {
	case "chat_message":
		message, err := handleChatMessage(rawEvent, server.Db)
		if err != nil {
			return err
		}
		server.Channels.Messages <- message

	case "direct_message":

		// server.Channels.Dm <- rawEvent

	case "user_login":
		server.Channels.Notify <- rawEvent

	// ? Non forwarded events - server only
	case "account_create":
		newAccount, err := handleAccountCreate(rawEvent, server.Db)
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
			logger.Err("ROUTER", "Error serializing")
		}

		rawEvent.Data = serializedData

		server.Channels.Accounts <- rawEvent

	case "user_logout":
		server.Channels.Notify <- rawEvent

	case "user_update":
		server.Channels.Notify <- rawEvent

	case "chat_create":
		server.Channels.Notify <- rawEvent

	case "chat_delete":
		server.Channels.Notify <- rawEvent

	case "chat_join":
		server.Channels.Notify <- rawEvent

	default:
		logger.Warn("HNDL", "Payload Handler received an invalid event type")
		return fmt.Errorf("invalid_request")
	}
	return nil
}

func handleChatMessage(rawEvent *RawEvent, db *gorm.DB) (*Message, error) {
	message := &Message{UserID: rawEvent.UserID}

	if err := message.Deserialize(rawEvent.Data); err != nil {
		logger.Err("HNDL", "Error Deserializing Chat Message")
		return nil, err
	}

	if err := message.SaveToDb(db); err != nil {
		logger.Err("HNDL", "Error saving to DB Chat Message")
		return nil, err
	}

	message.User = &User{ID: message.UserID}
	if err := message.User.GetInfo(db); err != nil {
		logger.Err("HNDL", fmt.Sprintf("Error getting user info ID: %d\n", message.UserID))
		return nil, err
	}

	message.Chat = &Chat{ID: message.ChatID}

	if _, err := message.Chat.GetInfo(db); err != nil {
		logger.Err("HNDL", fmt.Sprintf("Error getting message chat ID: %d\n", message.UserID))
		return nil, err
	}

	return message, nil

}

func handleAccountCreate(rawEvent *RawEvent, db *gorm.DB) (*Account, error) {
	account := &Account{}

	if err := account.Deserialize(rawEvent.Data); err != nil {
		logger.Err("HNDL", "Error Deserializing User")
		return nil, err
	}

	if err := account.Create(db); err != nil {
		logger.Err("HNDL", fmt.Sprintf("Error Creating Account: %d\n", account.ID))
		return nil, err
	}

	logger.Info("HNDL", fmt.Sprintf("Account Created: %d\n", account.ID))
	return account, nil
}

func handleUserCreate(rawEvent *RawEvent, db *gorm.DB, account *Account) (*User, error) {
	user := &User{Account: account}

	if err := user.Deserialize(rawEvent.Data); err != nil {
		logger.Err("HNDL", "Error Deserializing User")
		return nil, err
	}

	if err := user.Create(db); err != nil {
		logger.Err("HNDL", fmt.Sprintf("Error Creating User: %d\n", user.ID))
		return nil, err
	}

	logger.Info("HNDL", fmt.Sprintf("User Created: %d\n", user.ID))
	return user, nil
}
