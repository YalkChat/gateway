package handlers

import (
	"fmt"
	"log"
	"yalk/database"
	"yalk/database/models"
	"yalk/newchat/event"
	"yalk/newchat/event/types"
)

type NewMessageHandler struct{}

func (h NewMessageHandler) HandleEvent(ctx *event.HandlerContext, e event.Event) error {

	// Step 1: Convert e to NewMessageEvent type
	newMessageEvent, ok := e.(types.NewMessageEvent)
	if !ok {
		return fmt.Errorf("invalid event type:", e.Type())
	}

	// TODO: Evaluate if the logic of authentication and validation can be a Higher order function, or plug-n-play, whatever it's called
	// TODO: Evaluate if event to handle should be in struct properties or as function argument
	if err := validateMessageCreate(newMessageEvent); err != nil {
		log.Printf("Validation failed: %v", err)
		return err
	}

	// Step 2: Make Database type
	if err := database.AddMessage(ctx.DB, &models.Message{}); err != nil {
		return err
	}
	// Step 2: Database Operation
	if err := database.AddMessage(ctx.DB, &models.Message{}); err != nil {
		log.Printf("Database operation failed: %v", err)
		return err
	}

	// TODO: Change this placeholder
	if err := ctx.SendToChat("1", newMessageEvent); err != nil {

	}

	log.Printf("Handling MessageCreate event: %s", "placeholder")
	return nil
}

func validateMessageCreate(eventData types.NewMessageEvent) error {
	// Validate the message content, ChatID, SenderID, etc.
	// Return an error if validation fails
	return nil
}
