package initialize

import (
	"yalk/chat"

	"gorm.io/gorm"
)

func createAdminAccount(db *gorm.DB) (*chat.Account, error) {
	// ! Hash for default admin's "admin" password in BCrypt, it will not be this and
	// ! not be set this way.
	adminAccount := &chat.Account{
		Email:    "admin@example.com",
		Username: "admin",
		Password: "$2a$14$QuxLu/0REKoTuZGcwZwX2eLsNKFrook.QMh/Esd8d4FocaE2sKHca",
		Verified: true}

	if err := adminAccount.Create(db); err != nil {
		return nil, err
	}
	return adminAccount, nil
}

func createAdminUser(db *gorm.DB, adminAccount *chat.Account) (*chat.User, error) {
	adminUser := &chat.User{
		Account:       adminAccount,
		DisplayedName: "Admin",
		AvatarUrl:     "/default.png",
		StatusName:    "offline"}

	if err := adminUser.Create(db); err != nil {
		return nil, err
	}
	return adminUser, nil
}
