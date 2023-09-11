package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
	"yalk/chat/models/events"
)

// TODO: ChatType here has something wrong, I'm not sure why but it's wrong.
// TODO: Missing method in DatabaseOperations
func createMainChannel(dbConn database.DatabaseOperations) error {
	mainChat := &db.Chat{
		Name:      "Main",
		ChatType:  chatType,
		CreatedBy: adminUser,
		Users:     []*events.User{adminUser}}

	return mainChat, nil
}
