package server

import "yalk/chat/models/events"

func (s *serverImpl) GetUserByID(userID uint) (*events.User, error) {
	dbUser, err := s.db.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	// TODO: Finish implement
	user := events.User{ID: dbUser.ID}
	return &user, nil
}

func (s *serverImpl) GetUserByUsername(username string) (*events.User, error) {
	dbUser, err := s.db.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	// Add handle empty fields
	user := &events.User{ID: dbUser.ID}
	return user, nil
}
