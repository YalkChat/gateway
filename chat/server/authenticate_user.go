package server

import (
	"yalk/chat/models/events"
	"yalk/errors"

	"golang.org/x/crypto/bcrypt"
)

// Change return to user event if I need more info basides the user id
func (s *serverImpl) AuthenticateUser(userLogin events.UserLogin) (uint, error) {
	dbUser, err := s.db.GetUserByUsername(userLogin.Username)
	if err != nil {
		return 0, err
	}

	// Validate the password
	if !validatePassword(dbUser.Password, userLogin.Password) {
		return 0, errors.ErrAuthInvalid
	}

	return dbUser.ID, nil
}

func validatePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
