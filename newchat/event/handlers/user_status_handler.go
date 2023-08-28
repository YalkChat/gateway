package handlers

import (
	"yalk/newchat/event"

	"gorm.io/gorm"
)

func HandleUserStatusEvent(db *gorm.DB) event.Handler {
	return func(e event.Event) error {
		// Handle user status event
	}
}
