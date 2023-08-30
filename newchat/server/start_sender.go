package server

import (
	"yalk/newchat/client"
)

func (s *serverImpl) StartSender(c client.Client, quit chan struct{}) {
	for {
		select {
		case <-quit:
			// Handle cleanup if needed
			return
		default:
		}
	}

	// for event := range outgoingEvents {
	// 	if err := c.SendEvent(event); err != nil {
	// 		// handle or log the error
	// 	}
	// }
}
