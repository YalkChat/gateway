package events

type UserCreationEvent struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
