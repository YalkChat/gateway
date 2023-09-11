package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
)

// TODO: Needs to be adapter and methods defined in DatabaseOperations if necessary
func createBotUser(dbConn database.DatabaseOperations) error {
	botUser := &db.User{
		DisplayedName: "Bot",
		AvatarUrl:     "/bot.png",
		StatusID:      "bot"}

	if err := dbConn.NewUser(botUser); err != nil {
		return err
	}
	return nil
}
