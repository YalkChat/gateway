package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
)

func saveInitialSettings(dbCon database.DatabaseOperations) error {
	serverSettings := &db.ServerSettings{IsInitialized: true}
	if err := dbCon.SaveServerSettings(serverSettings); err != nil {
		return err
	}
	return nil
}
