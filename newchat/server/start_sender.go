package server

import (
	"yalk/newchat/client"
	"yalk/newchat/event"
)

func (s *serverImpl) StartSender(c client.Client, outgoingEvents chan event.Event) {
	for event := range outgoingEvents {
		if err := c.SendEvent(event); err != nil {
			// handle or log the error
		}
	}
}
