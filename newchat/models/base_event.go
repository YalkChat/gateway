package models

import "encoding/json"

type BaseEvent struct {
	Opcode   string          `gorm:"-" json:"opcode"`
	Data     json.RawMessage `gorm:"-" json:"data"`
	ClientID string          `gorm:"-" json:"clientID"`
	Type     string          `gorm:"-" json:"type"`
}
