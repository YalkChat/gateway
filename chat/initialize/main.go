package initialize

import (
	"fmt"
	"log"
	"yalk/chat/database"
	"yalk/chat/util"
)

// TODO: Break down further
func InitializeApp(dbConn database.DatabaseOperations) error {

	if isInitialized := checkIsInitialized(dbConn); isInitialized {
		return nil
	}

	err := createBotUser(dbConn)
	if err != nil {
		fmt.Printf("Can't create bot user, error: %v", err)
		return err
	}
	log.Printf("Created bot user")

	adminUser, err := createAdminUser(dbConn)
	if err != nil {
		fmt.Printf("Can't create admin profile, error: %v", err.Error())
		return err
	}
	log.Printf("Created admin user")

	// Convert db.User to events.User
	eventAdminUser := util.ConvertDBUserToEventUser(adminUser)

	// TODO: This needs to be moved to events.ChatType, how to do it?
	channelType, err := createChannelType(dbConn)
	if err != nil {
		fmt.Printf("Can't create chat type, error: %v", err)
		return err
	}
	log.Printf("Created channel types")

	// Convert db.ChannelType to events.User
	eventChannelType := util.ConvertDbChatTypeToEventChatType(channelType)

	err = createMainChannel(dbConn, eventChannelType, eventAdminUser)
	if err != nil {
		fmt.Printf("Can't create main chat, error: %v", err)
		return err
	}
	log.Printf("Created main channel")

	if err = saveInitialSettings(dbConn); err != nil {
		fmt.Printf("Can't save initial settings, error: %v", err)
		return err
	}
	log.Print("Created initial settings")

	return nil
}
