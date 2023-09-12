package initialize

import (
	"yalk/chat/database"
)

func checkIsInitialized(dbConn database.DatabaseOperations) bool {
	isInitialized, err := dbConn.IsServerInitialized()
	if err != nil {
		return false
	}
	return isInitialized
}
