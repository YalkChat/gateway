package database

import "yalk/newchat/models/db"

// TODO: Evaluate strings arguments used
// TODO: We could pass the type instance with the ID property set and use .Find() on it
// TODO: Decide and choose if these operations must have the db model passed as arg or a string based on the example above.
// TODO: Also need to decide the returns of these functions
type DatabaseOperations interface {
	SaveMessage(*db.Message) error
	GetMessage(string) (*db.Message, error)
	GetUsers(string) ([]string, error)
	NewUser(*db.User) (*db.User, error)
}
