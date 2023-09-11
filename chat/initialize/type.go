package initialize

import (
	"yalk/chat/database"

	"yalk/chat/models/db"
)

func createChannelType(dbConn database.DatabaseOperations) (*db.ChatType, error) {
	chatType := &db.ChatType{Name: "channel"}
	if err := dbConn.NewChatType(chatType); err != nil {
		return nil, err
	}
	return chatType, nil
}
