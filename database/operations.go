package database

import (
	"yalk/database/models"

	"gorm.io/gorm"
)

// AddMessage adds a new message to the database
func AddMessage(db *gorm.DB, message *models.Message) error {
	return db.Create(message).Error
}

func GetMessage(db *gorm.DB, chatID string) ([]models.Message, error) {
	var message []models.Message
	err := db.Where("chat_id = ?", chatID).Find(&message).Error
	return message, err
}

func RegisterClient(db *gorm.DB, client *models.Client) error {
	return db.Create(client).Error
}

func GetUsers(db *gorm.DB, chatID string) ([]string, error) {
	var chat models.Chat
	result := db.Preload("Users").First(&chat, "id = ?", chatID)
	if result.Error != nil {
		return nil, result.Error
	}

	var clientIDs []string
	for _, client := range chat.Clients {
		clientIDs = append(clientIDs, client.ID)
	}

	return clientIDs, nil
}

// func SaveMessage(db *gorm.DB, message *models.Message) error {

// }
