package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/events"
)

func createAdminUser(dbConn database.DatabaseOperations) error {
	adminUser := &events.UserCreationEvent{
		Email:     "admin@example.com",
		Password:  "$2a$14$QuxLu/0REKoTuZGcwZwX2eLsNKFrook.QMh/Esd8d4FocaE2sKHca",
		AvatarUrl: "/default.png",
		StatusID:  "offline"}

	if _, err := dbConn.NewUserWithPassword(adminUser); err != nil {
		return err
	}
	return nil
}
