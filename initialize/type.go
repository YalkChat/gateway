package initialize

import (
	"yalk/chat"

	"gorm.io/gorm"
)

func createChannelType(db *gorm.DB) (*chat.ChatType, error) {
	chatType := &chat.ChatType{Type: "channel"}

	tx := db.Create(chatType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return chatType, nil
}
