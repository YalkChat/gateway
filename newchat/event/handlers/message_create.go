package handlers

import (
	"yalk/newchat/event/types"

	"gorm.io/gorm"
)

func HandleMessageCreate(db *gorm.DB) func(e types.MessageCreateEvent) error {
	return func(e types.MessageCreateEvent) error {

		// Step 1: Validation
		if err := validateMessageCreate(e); err != nil {
			return err
		}

		// Step 2: Database Operation
		// messageID, err := database.AddMessage(db, message)
		// log.Printf("Handling MessageCreate event: %v", e)

		return nil
	}
}

func validateMessageCreate(eventData types.MessageCreateEvent) error {
	//  Valudate the message content, ChatID, SenderID, etc.
	// Return an error if validation fails
	return nil
}
