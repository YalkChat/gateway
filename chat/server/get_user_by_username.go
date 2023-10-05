package server

import (
	"yalk/chat/models/db"
)

func (s *serverImpl) GetUserByUsername(username string) (*db.User, error) {
	return s.db.GetUserByUsername(username)
}
