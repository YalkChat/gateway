package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
)

// TODO: ChatType here has something wrong, I'm not sure why but it's wrong.
// TODO: Needs to be adapter and methods defined in DatabaseOperations if necessary
func createMainChannel(dbConn database.DatabaseOperations, chatType *db.ChatType, adminUser *db.User) error {
	mainChat := &db.Chat{
		Name:     "Main",
		ChatType: chatType,
		// CreatedBy: adminUser,
		Users: []*db.User{adminUser},
	}
	if err := dbConn.NewChat(mainChat); err != nil {
		return err
	}
	return nil
}
