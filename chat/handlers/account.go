package handlers

import (
	"log"
	"yalk/chat/events"
	"yalk/chat/models"

	"gorm.io/gorm"
)

func HandleAccount(rawEvent *events.RawEvent, db *gorm.DB) (*models.Account, error) {
	account := &models.Account{}

	if err := account.Deserialize(rawEvent.Data); err != nil {
		log.Printf("Error Deserializing Account")
		return nil, err
	}

	if err := account.Create(db); err != nil {
		log.Printf("Error Creating Account: %d\n", account.ID)
		return nil, err
	}

	log.Printf("Account Created: %d\n", account.ID)
	return account, nil
}
