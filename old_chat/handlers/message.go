package handlers

import (
	"log"
	"yalk/chat/models"

	"gorm.io/gorm"
)

func HandleMessage(rawEvent *models.RawEvent, db *gorm.DB) (*models.Message, error) {
	var user *models.User
	var message *models.Message
	var chat *models.Chat

	switch rawEvent.Action {
	case "send":
		user = &models.User{ID: rawEvent.UserID}
		if err := user.GetInfo(db); err != nil {
			log.Printf("Error getting user info ID: %d\n", rawEvent.UserID)
			return nil, err

		}

		message = &models.Message{UserID: rawEvent.UserID}
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
