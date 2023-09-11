package events

import "time"

// TODO: Very important To-Do: it is possible that this type will need to
// TODO: become NewMessage, evaluate why it would be needed, the doubt is the
// TODO: additional informations that the event could need to carry but for now
// TODO: none came to my mind, and also it looks suspicious that it's almost
// TODO: identical to it's database models counterpart.

type Message struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ChatID    uint      `json:"chat_id"`
	UserID    uint      `json:"client_id"`
	Content   string    `json:"content"`
}
