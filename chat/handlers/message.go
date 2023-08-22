package handlers

import (
	"log"
	"yalk/chat/chatmodels"
	"yalk/database/dbmodels"

	"gorm.io/gorm"
)

func HandleMessage(rawEvent *chatmodels.RawEvent, db *gorm.DB) (*dbmodels.Message, error) {
	var user *dbmodels.User
	var message *dbmodels.Message
	var chat *dbmodels.Chat

	switch rawEvent.Action {
	case "send":
		user = &dbmodels.User{ID: rawEvent.UserID}
		if err := user.GetInfo(db); err != nil {
			log.Printf("Error getting user info ID: %d\n", rawEvent.UserID)
			return nil, err

		}

		message = &dbmodels.Message{UserID: rawEvent.UserID}
		if err := message.Deserialize(rawEvent.Data); err != nil {
			log.Printf("Error Deserializing Chat Message")
			return nil, err
		}

		// chat = &models.Chat{ID: message.ChatID}
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
