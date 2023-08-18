package main

import (
	"fmt"
	"log"
	"yalk/chat"
	"yalk/logger"

	"gorm.io/gorm"
)

func initializeApp(db *gorm.DB) error {

	// Check if already initialized
	if isInitialized := checkIsInitialized(db); isInitialized {
		return nil
	}

	// Creating bot account
	botAccount, err := createBotAccount(db)
	if err != nil {
		fmt.Printf("Can't create bot account, error: %v", err)
		return err
	}
	logger.Info("CORE", fmt.Sprintf("Created bot account: %v", botAccount.Username))

	// Creating bot user
	botUser, err := createBotUser(db, botAccount)
	if err != nil {
		fmt.Printf("Can't create bot user, error: %v", err)
		return err
	}
	log.Printf("Created bot account: %v", botUser.DisplayedName)

	// Creating initial admin account
	adminAccount, err := createAdminAccount(db)
	if err != nil {
		fmt.Printf("Can't create admin credentials: %v", err.Error())
		return err
	}
	log.Printf("Created admin credentials: %v", adminAccount.Username)

	// Creating initial admin user
	adminUser, err := createAdminUser(db, adminAccount)
	if err != nil {
		fmt.Printf("Can't create admin profile, error: %v", err.Error())
		return err
	}
	log.Printf("Created admin user: %v", adminUser.DisplayedName)

	chatType, err := createChannelType(db)
	if err != nil {
		fmt.Printf("Can't create chat type, error: %v", err)
		return err
	}
	log.Printf("Created channel types: %v", chatType.Type)

	mainChat, err := createMainChannel(db, adminUser, chatType)
	if err != nil {
		fmt.Printf("Can't create main chat, error: %v", err)
		return err
	}
	log.Printf("Created main channel: %v", mainChat.Name)

	if err = saveInitialSettings(db); err != nil {
		fmt.Printf("Can't save initial settings, error: %v", err)
		return err
	}
	log.Print("Created initial settings")

	return nil
}

func checkIsInitialized(db *gorm.DB) bool {
	var serverSettings *chat.ServerSettings
	tx := db.Select("is_initialized").First(&serverSettings, "is_initialized = true")
	return tx.Error == nil
}

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

func createBotAccount(db *gorm.DB) (*chat.Account, error) {
	botAccount := &chat.Account{
		Email:    "invalid@example.com",
		Username: "bot",
		Password: "none",
		Verified: false}

	if err := botAccount.Create(db); err != nil {
		return nil, err
	}
	return botAccount, nil
}

func createBotUser(db *gorm.DB, botAccount *chat.Account) (*chat.User, error) {
	botUser := &chat.User{
		DisplayedName: "Bot",
		AvatarUrl:     "/bot.png",
		Account:       botAccount,
		StatusName:    "bot"}

	if err := botUser.Create(db); err != nil {
		return nil, err
	}
	return botUser, nil
}

func createChannelType(db *gorm.DB) (*chat.ChatType, error) {
	chatType := &chat.ChatType{Type: "channel"}

	tx := db.Create(chatType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return chatType, nil
}

// TODO: ChatType here has something wrong, I'm not sure why but it's wrong.
func createMainChannel(db *gorm.DB, adminUser *chat.User, chatType *chat.ChatType) (*chat.Chat, error) {
	mainChat := &chat.Chat{
		Name:      "Main",
		ChatType:  chatType,
		CreatedBy: adminUser,
		Users:     []*chat.User{adminUser}}

	tx := db.Create(mainChat)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mainChat, nil
}

func saveInitialSettings(db *gorm.DB) error {
	serverSettings := &chat.ServerSettings{IsInitialized: true}
	err := serverSettings.Create(db)
	if err != nil {
		return err
	}
	return nil
}
