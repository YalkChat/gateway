package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
	"yalk/chat/models/events"
)

func createAdminUser(dbConn database.DatabaseOperations) (*db.User, error) {
	adminUser := &events.UserCreationEvent{
		Email:     "admin@example.com",
		Password:  "$2a$14$QuxLu/0REKoTuZGcwZwX2eLsNKFrook.QMh/Esd8d4FocaE2sKHca",
		AvatarUrl: "/default.png",
		StatusID:  "offline"}

	dbAdminUser, err := dbConn.NewUserWithPassword(adminUser)
	if err != nil {
		return nil, err
	}
	return dbAdminUser, nil
}
