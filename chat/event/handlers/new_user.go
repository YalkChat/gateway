package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"yalk/chat/database"
	"yalk/chat/event"
	"yalk/chat/models/db"
	"yalk/chat/models/events"
)

type NewUserHandler struct{}

func (h NewUserHandler) HandleEvent(ctx *event.HandlerContext, baseEvent *events.BaseEvent) error {
	newUser, err := parseUserCreationEvent(baseEvent)
	if err != nil {
		log.Printf("Error parsing user: %v", err)
		return err
	}
	// TODO: Check if this actually updates the struct with the new ID
	newDbUser, err := saveNewUserToDb(newUser, ctx.DB)
	if err != nil {
		return fmt.Errorf("error creating new user: %v", err)
	}

	// Create a new User object from the database user
	newUserEvent := &events.User{
		ID:    newDbUser.ID,
		Email: newDbUser.Email,
		// Add other fields as needed
	}

	// Serialize the new user event
	serializedData, err := ctx.Serializer.Serialize(newUserEvent)
	if err != nil {
		return fmt.Errorf("error serializing new user: %v", err)
	}

	newBaseEvent := &events.BaseEvent{
		Opcode:   "NewUser", //TODO: This is not the right opcode
		Data:     serializedData,
		ClientID: newDbUser.ID,
		Type:     "placeholder",
	}

	ctx.SendAll(newBaseEvent)

	return nil
}

func parseUserCreationEvent(baseEvent *events.BaseEvent) (*events.UserCreationEvent, error) {
	var newUserCreationEvent *events.UserCreationEvent
	if err := json.Unmarshal(baseEvent.Data, &newUserCreationEvent); err != nil {
		return nil, fmt.Errorf("error parsing new user: %v", err)
	}
	return newUserCreationEvent, nil
}

// TODO: check whether we can have a better way to return the ID
// TODO: finish implementation
// TODO: Could use github.com/ulule/deepcopier
func saveNewUserToDb(newUser *events.UserCreationEvent, database database.DatabaseOperations) (*db.User, error) {
	// We need to get the new ID back
	dbNewUser, err := database.NewUserWithPassword(newUser)
	if err != nil {
		return nil, fmt.Errorf("error saving new user: %v", err)
	}
	return dbNewUser, nil
}
