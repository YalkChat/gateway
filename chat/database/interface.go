package database

import (
	"yalk/chat/models/db"
	"yalk/chat/models/events"
)

// TODO: Evaluate strings arguments used
// TODO: We could pass the type instance with the ID property set and use .Find() on it
// TODO: Decide and choose if these operations must have the db model passed as arg or a string based on the example above.
// TODO: Also need to decide the returns of these functions

// TODO: Move to use the events models instead of db models
type DatabaseOperations interface {
	SaveMessage(*events.Message) (*db.Message, error)
	GetMessage(uint) (*db.Message, error)
	GetUsersByChatId(uint) ([]*db.User, error)
	NewUser(*events.User) (*db.User, error)
	NewChat(*events.Chat) (*db.Chat, error)
	NewChatType(*events.ChatType) (*db.ChatType, error)
	NewUserWithPassword(*events.UserCreationEvent) (*db.User, error)
	IsServerInitialized() (bool, error)
	SaveServerSettings(*events.ServerSettings) error
}
