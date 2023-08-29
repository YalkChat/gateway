package event

import (
	"encoding/json"
	"yalk/newchat/models"

	"gorm.io/gorm"
)

type Event interface {
	Type() string
	Data() json.RawMessage
	ClientID() string
	// Other methods as needed
}

type HandlerContext struct {
	DB         *gorm.DB
	SendToChat func(models.Message) error
}

// Handler defines the methods that any event handler must implement
// TODO: I must chose whether I want to keep the DB here, or use something else
type Handler interface {
	HandleEvent(*HandlerContext, *models.BaseEvent) error
}
