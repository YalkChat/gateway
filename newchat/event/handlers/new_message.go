package handlers

import (
	"log"
	"yalk/database"
	"yalk/database/models"

	"gorm.io/gorm"
)

type NewMessageHandler struct{}

func (e NewMessageHandler) HandleEvent(db *gorm.DB, event NewMessageHandler) error {
	// Step 1: Validation
	// TODO: Evaluate if the logic of authentication and validation can be a Higher order function, or plug-n-play, whatever it's called
	// TODO: Evaluate if event to handle should be in struct properties or as function argument
	if err := validateMessageCreate(event); err != nil {
		log.Printf("Validation failed: %v", err)
		return err
	}

	// Step 2: Make Database type

	if err := database.AddMessage(db, &models.Message{}); err != nil {
		return err
	}
	// Step 2: Database Operation
	if err := database.AddMessage(db, &models.Message{}); err != nil {
		log.Printf("Database operation failed: %v", err)
		return err
	}

	log.Printf("Handling MessageCreate event: %v", e)
	return nil
}

func validateMessageCreate(eventData NewMessageHandler) error {
	// Validate the message content, ChatID, SenderID, etc.
	// Return an error if validation fails
	return nil
}
