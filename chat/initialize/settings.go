package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/events"
)

func saveInitialSettings(dbCon database.DatabaseOperations) error {
	serverSettings := &events.ServerSettings{IsInitialized: true}
	if err := dbCon.SaveServerSettings(serverSettings); err != nil {
		return err
	}
	return nil
}
