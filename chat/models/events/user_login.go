package events

// LoginPayload represents the payload for a user login event.
// TODO: Can I remove the password from the User type then?
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
