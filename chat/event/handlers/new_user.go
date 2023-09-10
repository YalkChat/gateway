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
	newUser, err := parseUser(baseEvent)
	if err != nil {
		log.Printf("Error parsing user: %v", err)
		return err
	}
	// TODO: Check if this actually updates the struct with the new ID
	newDbUser, err := saveNewUserToDb(newUser, ctx.DB)
	if err != nil {
		return fmt.Errorf("error creating new user: %v", err)
	}

	ctx.SendToAll(newDbUser)

	return nil
}

func parseUser(baseEvent *events.BaseEvent) (*events.User, error) {
	var newUser *events.User
	if err := json.Unmarshal(baseEvent.Data, &newUser); err != nil {
		return nil, fmt.Errorf("error parsing new account: %v", err)
	}
	return newUser, nil
}

// TODO: check whether we can have a better way to return the ID
// TODO: finish implementation
// TODO: Could use github.com/ulule/deepcopier
func saveNewUserToDb(newUser *events.User, database database.DatabaseOperations) (*db.User, error) {
	dbUser := &db.User{Email: newUser.Email}

	// We need to get the new ID back
	dbUserWithID, err := database.NewUser(dbUser)
	if err != nil {
		return nil, fmt.Errorf("error saving new user: %v", err)
	}
	return dbUserWithID, nil
}
