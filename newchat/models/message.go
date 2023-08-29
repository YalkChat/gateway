package models

import "time"

type Message struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ChatID    string    `json:"chat_id"`
	ClientID  string    `json:"client_id"`
	Content   string    `json:"content"`
}
