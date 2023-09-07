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

type NewMessageHandler struct{}

func (h NewMessageHandler) HandleEvent(ctx *event.HandlerContext, baseEvent *events.BaseEvent) error {

	// Step 1: Parse baseEvent to Message type
	newMessage, err := parseMessage(baseEvent)
	if err != nil {
		log.Printf("Error parsing message: %v", err)
		return err
	}

	// TODO: Evaluate if the logic of authentication and validation can be a Higher order function, or plug-n-play, whatever it's called
	// TODO: Evaluate if event to handle should be in struct properties or as function argument
	// Step 2: Validate message
	if err := validateMessageCreate(newMessage); err != nil {
		log.Printf("Error validating message: %v", err)
		return err
	}

	// Step 3: Database Operation
	if err := saveToDatabase(newMessage, ctx.DB); err != nil {
		log.Printf("Error sending message to chat: %v", err)
		return err
	}

	// Step 4: Send to other clients
	if err := ctx.SendToChat(newMessage); err != nil {
		log.Printf("Error sending message to chat: %v", err)
		return err
	}

	log.Printf("Handling MessageCreate event: %s", "placeholder")
	return nil
}

func saveToDatabase(newMessage *events.Message, database database.DatabaseOperations) error {
	dbMessage := &db.Message{
		ChatID:   newMessage.ChatID,
		ClientID: newMessage.ClientID,
		Content:  newMessage.Content,
	}

	if err := database.SaveMessage(dbMessage); err != nil {
		log.Printf("Database operation failed: %v", err)
		return err
	}
	return nil
}

func parseMessage(baseEvent *events.BaseEvent) (*events.Message, error) {
	var newMessage *events.Message
	if err := json.Unmarshal(baseEvent.Data, &newMessage); err != nil {
		return nil, fmt.Errorf("error parsing message: %v", err)
	}
	return newMessage, nil
}

func validateMessageCreate(eventData *events.Message) error {
	// Validate the message content, ChatID, SenderID, etc.
	// Return an error if validation fails
	// TODO: Write implementation
	// if err != nil {
	// 	return fmt.Errorf("Validation failed: %v", err)
	// }
	return nil
}
