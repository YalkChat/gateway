package initialize

import (
	"yalk/chat/database"

	"yalk/chat/models/db"
	"yalk/chat/models/events"
)

func createChannelType(dbConn database.DatabaseOperations) (*db.ChatType, error) {
	chatType := &events.ChatType{Name: "channel"}
	DbChatType, err := dbConn.NewChatType(chatType)
	if err != nil {
		return nil, err
	}
	return DbChatType, nil
}
