package initialize

import (
	"yalk/chat/database"
	"yalk/chat/models/db"
)

func checkIsInitialized(dbConn database.DatabaseOperations) bool {
	var serverSettings *db.ServerSettings
	isInitialized, err := dbConn.IsServerInitialized()
	if err != nil {
		return false
	}
	return isInitialized
}
