package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
)

// TODO: Missing method in DatabaseOperations
func createBotUser(conn database.DatabaseOperations) error {
	botUser := &db.User{
		DisplayedName: "Bot",
		AvatarUrl:     "/bot.png",
		StatusID:      "bot"}

	return nil
}
