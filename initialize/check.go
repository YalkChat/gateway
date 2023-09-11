package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
)

// TODO: Missing method in DatabaseOperations
func checkIsInitialized(dbConn database.DatabaseOperations) bool {
	var serverSettings *db.ServerSettings
	tx := db.Select("is_initialized").First(&serverSettings, "is_initialized = true")
	return tx.Error == nil
}
