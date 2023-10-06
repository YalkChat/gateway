package server

import (
	"yalk/chat/models/db"
	"yalk/errors"

	"golang.org/x/crypto/bcrypt"
)

// Change return to user event if I need more info basides the user id
func (s *serverImpl) AuthenticateUser(user *db.User) (uint, error) {
	dbUser, err := s.db.GetUserByUsername(user.Username)
	if err != nil {
		return 0, err
	}

	// Validate the password
	if !validatePassword(dbUser.Password, user.Password) {
		return 0, errors.ErrAuthInvalid
	}

	return dbUser.ID, nil
}

func validatePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
