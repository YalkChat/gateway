package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
)

// TODO: Missing method in DatabaseOperations
func saveInitialSettings(dbCon database.DatabaseOperations) error {
	serverSettings := &db.ServerSettings{IsInitialized: true}

	return nil
}
