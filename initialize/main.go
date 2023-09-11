package initialize

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func InitializeApp(db *gorm.DB) error {

	if isInitialized := checkIsInitialized(db); isInitialized {
		return nil
	}

	botUser, err := createBotUser(db, botAccount)
	if err != nil {
		fmt.Printf("Can't create bot user, error: %v", err)
		return err
	}
	log.Printf("Created bot user: %v", botAccount.DisplayedName)

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
