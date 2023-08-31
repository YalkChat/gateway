package database

import (
	"yalk/database/models"

	"gorm.io/gorm"
)

type DatabaseImpl struct {
	conn *gorm.DB
}

func NewDatabase(conn *gorm.DB) *DatabaseImpl {
	return &DatabaseImpl{
		conn: conn,
	}
}

func (db *DatabaseImpl) SaveMessage(message *models.Message) error {
	return db.conn.Create(message).Error

}

func (db *DatabaseImpl) GetMessage(messageID string) (*models.Message, error) {
	var message *models.Message
	if err := db.conn.Where("id = ?", messageID).First(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (db *DatabaseImpl) GetUsers(chatID string) ([]string, error) {
	var chat models.Chat
	result := db.conn.Preload("Users").First(&chat, "id = ?", chatID)
	if result.Error != nil {
		return nil, result.Error
	}

	var clientIDs []string
	for _, client := range chat.Clients {
		clientIDs = append(clientIDs, client.ID)
	}

	return clientIDs, nil
}

// TODO: Decide what this function should reeturn
func (db *DatabaseImpl) NewUser(newUser *models.User) (*models.User, error) {
	if err := db.conn.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}
