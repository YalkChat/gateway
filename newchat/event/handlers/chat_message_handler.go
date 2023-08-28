package handlers

import (
	"yalk/newchat/event"

	"gorm.io/gorm"
)

func HandleChatMessageEvent(db *gorm.DB) event.Handler {
	return func(e event.Event) error {
		// Handle chat message event
	}
}
