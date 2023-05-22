package chat

import (
	"fmt"

	"yalk/logger"
	"yalk/sessions"

	"gorm.io/gorm"
)

// * Handle incoming user payload and process it eventually
// * forwarding in the correct routine channels for other users to receive.
func (server *Server) HandlePayload(clientId uint, jsonEventMessage []byte) error {
	rawEvent := &RawEvent{}

	if err := rawEvent.Deserialize(jsonEventMessage); err != nil {
		logger.Err("HNDL", "Error unmarshaling RawEvent")
		return err
	}

	// server.SessionsManager.Validate()

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

	case "account_create":
		account, err := handleAccountCreate(rawEvent, server.Db)
		if err != nil {
			return err
		}
		server.Channels.Accounts <- account

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
	message, err := newMessage(rawEvent)
	if err != nil {
		logger.Err("HNDL", "Error creating Chat Message")
		return nil, err
	}

	if err := message.SaveToDb(db); err != nil {
		logger.Err("HNDL", "Error saving to DB Chat Message")
		return nil, err
	}

	message.User = &User{ID: message.UserID}
	message.User.GetInfo(db)
	if err != nil {
		logger.Err("HNDL", fmt.Sprintf("Error getting user info ID: %d\n", message.UserID))
		return nil, err
	}

	message.Chat = &Chat{ID: message.ChatID}
	message.Chat.GetInfo(db)
	if err != nil {
		logger.Err("HNDL", fmt.Sprintf("Error getting message chat ID: %d\n", message.UserID))
		return nil, err
	}

	// if err != nil {
	// 	logger.Err("PROFILE", fmt.Sprintf("Error getting Chat ID: %d\n", message.ChatID))
	// 	return nil, err
	// }
	return message, nil

}

func newMessage(rawEvent *RawEvent) (*Message, error) {
	message := &Message{}

	if err := message.Deserialize(rawEvent.Data); err != nil {
		logger.Err("HNDL", "Error Deserializing Chat Message")
		return nil, err
	}
	return message, nil
}

func handleAccountCreate(rawEvent *RawEvent, db *gorm.DB) (*sessions.Account, error) {
	account := &sessions.Account{}

	if err := account.Deserialize(rawEvent.Data); err != nil {
		logger.Err("HNDL", "Error Deserializing User")
		return nil, err
	}

	if err := account.Create(db); err != nil {
		logger.Err("HNDL", fmt.Sprintf("Error Creating Account: %d\n", account.ID))
		return nil, err
	}
	logger.Err("HNDL", "Error Deserializing User")
	logger.Info("HNDL", fmt.Sprintf("Account Created: %d\n", account.ID))
	return account, nil
}
