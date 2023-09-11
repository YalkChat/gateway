package database

import (
	"yalk/chat/models/db"
	"yalk/chat/models/events"

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

// TODO: Adapt to accept the
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

func (dbi *DatabaseImpl) GetUsers(chatID string) ([]uint, error) {
	var chat db.Chat
	result := dbi.conn.Preload("Users").Find(&chat, "id = ?", chatID)
	if result.Error != nil {
		return nil, result.Error
	}

	var userIDs []uint
	for _, user := range chat.Users {
		userIDs = append(userIDs, user.ID)
	}

	return userIDs, nil
}

// TODO: Decide what this function should reeturn
func (dbi *DatabaseImpl) NewUserWithPassword(newUser *events.UserCreationEvent) (*db.User, error) {
	dbNewUser := &db.User{Email: newUser.Email, Password: newUser.Password}
	if err := dbi.conn.Create(newUser).Error; err != nil {
		return nil, err
	}
	return dbNewUser, nil
}
