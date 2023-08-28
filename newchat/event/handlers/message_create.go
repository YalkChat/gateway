package handlers

import (
	"log"
	"yalk/database"
	"yalk/newchat/event/types"

	"gorm.io/gorm"
)

func HandleMessageCreate(db *gorm.DB) func(e types.MessageCreateEvent) error {
	return func(e types.MessageCreateEvent) error {
		// Step 1: Validation
		if err := validateMessageCreate(e); err != nil {
			log.Printf("Validation failed: %v", err)
			return err
		}

		// Step 2: Database Operation
		if err := database.AddMessage(db, e.Message); err != nil {
			log.Printf("Database operation failed: %v", err)
			return err
		}

		log.Printf("Handling MessageCreate event: %v", e)
		return nil
	}
}

func validateMessageCreate(eventData types.MessageCreateEvent) error {
	// Validate the message content, ChatID, SenderID, etc.
	// Return an error if validation fails
	return nil
}
