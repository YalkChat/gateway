package server

import (
	"yalk/chat/models/db"
)

func (s *serverImpl) GetUserByID(userID uint) (*db.User, error) {
	return s.db.GetUserByID(userID)
}
