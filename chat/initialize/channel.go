package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/events"
)

// TODO: ChatType here has something wrong, I'm not sure why but it's wrong.
// TODO: Needs to be adapter and methods defined in DatabaseOperations if necessary
func createMainChannel(dbConn database.DatabaseOperations, chatType *events.ChatType, adminUser *events.User) error {
	mainChat := &events.Chat{
		Name:      "Main",
		ChatType:  chatType,
		CreatedBy: adminUser,
		Users:     []*events.User{adminUser},
	}
	_, err := dbConn.NewChat(mainChat)
	if err != nil {
		return err
	}
	return nil
}
