package events

// LoginPayload represents the payload for a user login event.
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
