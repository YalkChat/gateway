package initialize

import (
	"yalk/database/dbmodels"

	"gorm.io/gorm"
)

func createChannelType(db *gorm.DB) (*dbmodels.ChatType, error) {
	chatType := &dbmodels.ChatType{Type: "channel"}

	tx := db.Create(chatType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return chatType, nil
}
