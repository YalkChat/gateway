package initialize

import (
	"yalk/chat/models"

	"gorm.io/gorm"
)

func createChannelType(db *gorm.DB) (*models.ChatType, error) {
	chatType := &models.ChatType{Type: "channel"}

	tx := db.Create(chatType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return chatType, nil
}
