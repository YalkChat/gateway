package handlers

import (
	"encoding/json"
	"log"
	"yalk/database"
	dbmodels "yalk/database/models"
	"yalk/newchat/event"

	servermodels "yalk/newchat/models"
)

type NewMessageHandler struct{}

func (h NewMessageHandler) HandleEvent(ctx *event.HandlerContext, baseEvent *servermodels.BaseEvent) error {
	var newMessage servermodels.Message
	json.Unmarshal(baseEvent.Data, &newMessage)
	// Step 1: Convert e to NewMessageEvent type
	// newMessageEvent, ok := e.(types.NewMessageEvent)
	// if !ok {
	// 	return fmt.Errorf("invalid event type:", e.Type())
	// }

	// TODO: Evaluate if the logic of authentication and validation can be a Higher order function, or plug-n-play, whatever it's called
	// TODO: Evaluate if event to handle should be in struct properties or as function argument
	if err := validateMessageCreate(newMessage); err != nil {
		log.Printf("Validation failed: %v", err)
		return err
	}

	dbmessage := &dbmodels.Message{
		ChatID:   newMessage.ChatID,
		ClientID: newMessage.ClientID,
		Content:  newMessage.Content,
	}

	// Step 2: Database Operation
	if err := database.AddMessage(ctx.DB, dbmessage); err != nil {
		log.Printf("Database operation failed: %v", err)
		return err
	}

	// TODO: Change this placeholder
	if err := ctx.SendToChat(newMessage); err != nil {

	}

	log.Printf("Handling MessageCreate event: %s", "placeholder")
	return nil
}

func validateMessageCreate(eventData servermodels.Message) error {
	// Validate the message content, ChatID, SenderID, etc.
	// Return an error if validation fails
	return nil
}
