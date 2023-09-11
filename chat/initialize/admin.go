package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/events"
)

func createAdminUser(dbConn database.DatabaseOperations) error {
	// TODO: How to set other properties with this struct?
	adminUser := &events.UserCreationEvent{
		Email:     "admin@example.com",
		Password:  "$2a$14$QuxLu/0REKoTuZGcwZwX2eLsNKFrook.QMh/Esd8d4FocaE2sKHca",
		AvatarUrl: "/default.png",
		StatusID:  "offline"}

	dbConn.NewUserWithPassword(adminUser)

	return nil
}
