package models

import "encoding/json"

type RawEvent struct {
	Type   string          `gorm:"-" json:"type"`
	Action string          `gorm:"-" json:"action"`
	UserID uint            `gorm:"-" json:"userId"`
	Data   json.RawMessage `gorm:"-" json:"data"`
}

func (event *RawEvent) Serialize() ([]byte, error) {
	return json.Marshal(event)
}

func (event *RawEvent) Deserialize(jsonEvent []byte) error {
	return json.Unmarshal(jsonEvent, event)
}
