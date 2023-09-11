package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
)

func createBotUser(conn database.DatabaseOperations) error {
	botUser := &db.User{
		DisplayedName: "Bot",
		AvatarUrl:     "/bot.png",
		StatusID:      "bot"}

	return nil
}
