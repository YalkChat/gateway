package database

import (
	"yalk/chat/models/db"

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

func (dbi *DatabaseImpl) SaveMessage(message *db.Message) error {
	return dbi.conn.Create(message).Error

}

func (dbi *DatabaseImpl) GetMessage(messageID string) (*db.Message, error) {
	var message *db.Message
	if err := dbi.conn.Where("id = ?", messageID).First(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (dbi *DatabaseImpl) GetUsers(chatID string) ([]string, error) {
	var chat db.Chat
	result := dbi.conn.Preload("Users").First(&chat, "id = ?", chatID)
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
func (dbi *DatabaseImpl) NewUser(newUser *db.User) (*db.User, error) {
	if err := dbi.conn.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}
