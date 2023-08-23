package chatmodels

import (
	"yalk/database/dbmodels"
)

type Event interface {
	Type() string
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
	SaveToDb() error
}

// TODO: Return &ServerMessageChannels
func MakeEventChannels() *EventChannels {
	return &EventChannels{
		Messages: make(chan *dbmodels.Message, 1),
		Accounts: make(chan *RawEvent, 1),
		Users:    make(chan *RawEvent, 1),
		Notify:   make(chan *RawEvent, 1),
		Cmd:      make(chan *RawEvent),
		Login:    make(chan *RawEvent),
		Logout:   make(chan *RawEvent),
	}
}
