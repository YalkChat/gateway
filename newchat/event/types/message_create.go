package types

type MessageCreateEvent struct {
	Content string `json:"content"`
	ChatID  string `json:"chat_id"`
}
