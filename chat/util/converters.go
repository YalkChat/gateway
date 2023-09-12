package util

import (
	"yalk/chat/models/db"
	"yalk/chat/models/events"
)

// ConvertDBUserToEventUser converts a database User model to an events User model
func ConvertDBUserToEventUser(dbUser *db.User) *events.User {
	return &events.User{
		ID:    dbUser.ID,
		Email: dbUser.Email,
		// Add other fields as needed
	}
}

// ConvertDbChatTypeToEventChatType converts a db.ChatType to an events.ChatType
func ConvertDbChatTypeToEventChatType(dbChatType *db.ChatType) *events.ChatType {
	return &events.ChatType{
		ID: dbChatType.ID,
		// ... other fields
	}
}
