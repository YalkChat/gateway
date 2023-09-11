package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models"
	"yalk/chat/models/db"
)

func createChannelType(dbConn database.DatabaseOperations) (*models.ChatType, error) {
	chatType := &db.ChatType{Name: "channel"}

	return chatType, nil
}
