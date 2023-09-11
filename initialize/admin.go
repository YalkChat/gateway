package initialize

import (
	"yalk/chat/models/db"
	"yalk/chat/models/events"

	"gorm.io/gorm"
)

// TODO: Finish implementation
func createAdminAccount(db *gorm.DB) (*events.User, error) {
	// ! Hash for default admin's "admin" password in BCrypt, it will not be this and
	// ! not be set this way.
	adminUser := &db.UserCreationEvent{
		Email:         "admin@example.com",
		DisplayedName: "admin"}

}

func createAdminUser(dbConn *gorm.DB, adminAccount *events.User) (*events.User, error) {
	adminUser := &db.User{
		DisplayedName: "Admin",
		Password:      "$2a$14$QuxLu/0REKoTuZGcwZwX2eLsNKFrook.QMh/Esd8d4FocaE2sKHca",

		AvatarUrl: "/default.png",
		StatusID:  "offline"}

	if err := adminUser.Create(dbConn); err != nil {
		return nil, err
	}
	return adminUser, nil
}
