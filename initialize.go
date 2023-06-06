package main

import (
	"fmt"
	"yalk/chat"
	"yalk/logger"

	"gorm.io/gorm"
)

func createAdmin(db *gorm.DB) (*chat.User, error) {
	// ! Hash for default admin's "admin" password in BCrypt, it will not be this and
	// ! not be set this way.
	adminAccountPwd := "$2a$14$QuxLu/0REKoTuZGcwZwX2eLsNKFrook.QMh/Esd8d4FocaE2sKHca"
	adminAccount := &chat.Account{Email: "admin@example.com", Username: "admin", Password: adminAccountPwd, Verified: true}
	err := adminAccount.Create(db)
	if err != nil {
		logger.Err("CORE", fmt.Sprintf("FATAL - Can't create admin credentials, error: %v", err))
		return nil, err
	}
	logger.Info("CORE", fmt.Sprintf("Created admin credentials: %v", adminAccount.Username))

	adminUser := &chat.User{Account: adminAccount, DisplayedName: "Admin", AvatarUrl: "/default.png"}
	err = adminUser.Create(db)
	if err != nil {
		logger.Err("CORE", fmt.Sprintf("FATAL - Can't create admin profile, error: %v", err.Error()))
		return nil, err
	}
	logger.Info("CORE", fmt.Sprintf("Created admin user: %v", adminUser.DisplayedName))
	return adminUser, nil
}

func checkIsInitialized(db *gorm.DB) bool {
	var serverSettings *chat.ServerSettings
	tx := db.Select("is_initialized").First(&serverSettings, "is_initialized = true")
	return tx.Error == nil
}

func createBotUser(db *gorm.DB) error {
	botAccountPwd := "none"
	botAccount := &chat.Account{Email: "invalid@example.com", Username: "bot", Password: botAccountPwd, Verified: false}
	err := botAccount.Create(db)
	if err != nil {
		logger.Err("CORE", fmt.Sprintf("FATAL - Can't create bot credentials, error: %v", err))
		return nil
	}
	logger.Info("CORE", fmt.Sprintf("Created bot credentials: %v", botAccount.Username))
	serverBot := &chat.User{DisplayedName: "Bot", AvatarUrl: "/bot.png", Account: botAccount}
	err = serverBot.Create(db)
	if err != nil {
		return err
	}
	return nil
}
