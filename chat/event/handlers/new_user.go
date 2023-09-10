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

	if err := saveToDatabase(newUser, ctx.DB); err != nil {
		log.Printf("Error sending message to chat: %v", err)
		return err
	}
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
func saveNewUserToDb(newUser *events.User, database database.DatabaseOperations) (*db.User, error) {
	dbUser := &db.User{}

	// We need to get the new ID back
	dbUserWithID, err := database.NewUser(dbUser)
	if err != nil {
		return nil, fmt.Errorf("error saving new user: %v", err)
	}
	return dbUserWithID, nil
}
