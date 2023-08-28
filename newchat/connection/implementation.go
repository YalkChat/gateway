package connection

import "nhooyr.io/websocket"

type connectionImpl struct {
	conn *websocket.Conn
}
