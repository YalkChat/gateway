package chatmodels

import (
	"yalk/database/dbmodels"
)

// TODO: Remove while destructuring the channel message send mechanism!
type EventChannels struct {
	Messages chan *dbmodels.Message // TODO Continue to test implementation of types as payload data
	Accounts chan *RawEvent
	Users    chan *RawEvent
	Notify   chan *RawEvent
	Cmd      chan *RawEvent
	Login    chan *RawEvent
	Logout   chan *RawEvent
}
