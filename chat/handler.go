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
			logger.Err("ROUTER", "Error serializing")
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
			logger.Err("ROUTER", "Error serializing")
		}

		rawEvent.Data = serializedData
		server.Channels.Users <- rawEvent

	default:
		logger.Warn("HNDL", "Payload Handler received an invalid event type")
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
			logger.Err("HNDL", fmt.Sprintf("Error getting user info ID: %d\n", rawEvent.UserID))
			return nil, err

		}

		message = &Message{UserID: rawEvent.UserID, User: user}
		if err := message.Deserialize(rawEvent.Data); err != nil {
			logger.Err("HNDL", "Error Deserializing Chat Message")
			return nil, err
		}

		chat = &Chat{ID: message.ChatID}
		if _, err := chat.GetInfo(db); err != nil {
			logger.Err("HNDL", fmt.Sprintf("Error getting message chat ID: %d\n", message.UserID))
			return nil, err
		}

		message.Chat = chat

		if err := message.SaveToDb(db); err != nil {
			logger.Err("HNDL", "Error saving to DB Chat Message")
			return nil, err
		}

	}
	return message, nil

	// message := &Message{UserID: rawEvent.UserID}

	// if err := message.Deserialize(rawEvent.Data); err != nil {
	// 	logger.Err("HNDL", "Error Deserializing Chat Message")
	// 	return nil, err
	// }

	// if err := message.SaveToDb(db); err != nil {
	// 	logger.Err("HNDL", "Error saving to DB Chat Message")
	// 	return nil, err
	// }

	// message.User = &User{ID: message.UserID}
	// if err := message.User.GetInfo(db); err != nil {
	// 	logger.Err("HNDL", fmt.Sprintf("Error getting user info ID: %d\n", message.UserID))
	// 	return nil, err
	// }

	// message.Chat = &Chat{ID: message.ChatID}

	// if _, err := message.Chat.GetInfo(db); err != nil {
	// 	logger.Err("HNDL", fmt.Sprintf("Error getting message chat ID: %d\n", message.UserID))
	// 	return nil, err
	// }

	// return message, nil

}

func handleAccount(rawEvent *RawEvent, db *gorm.DB) (*Account, error) {
	account := &Account{}

	if err := account.Deserialize(rawEvent.Data); err != nil {
		logger.Err("HNDL", "Error Deserializing Account")
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

func handleUser(rawEvent *RawEvent, db *gorm.DB) (*User, error) {
	var user = &User{ID: rawEvent.UserID}
	// var status = &Status{}

	if err := user.GetInfo(db); err != nil {
		logger.Err("HNDL", fmt.Sprintf("Error getting user info ID: %d\n", user.ID))
		return nil, err
	}

	switch rawEvent.Action {
	case "change_status":
		if err := user.Deserialize(rawEvent.Data); err != nil {
			logger.Err("HNDL", "Error Deserializing User")
			return nil, err
		}
		// user.ChangeStatus(db, user.StatusName)
		if err := user.SaveToDb(db); err != nil {
			logger.Err("HNDL", fmt.Sprintf("Error saving to DB User: %d\n", user.ID))
			return nil, err
		}
	}
	// user := &User{}

	// if err := user.Deserialize(rawEvent.Data); err != nil {
	// 	logger.Err("HNDL", "Error Deserializing User")
	// 	return nil, err
	// }

	// if err := user.GetInfo(db); err != nil {
	// 	logger.Err("HNDL", fmt.Sprintf("Error getting user info ID: %d\n", user.ID))
	// 	return nil, err
	// }

	// user.ChangeStatus(db, user.StatusName)
	// user.Status = rawEvent.Status

	// if err := user.SaveToDb(db); err != nil {
	// 	logger.Err("HNDL", fmt.Sprintf("Error saving to DB User: %d\n", user.ID))
	// 	return nil, err
	// }

	return user, nil
}
