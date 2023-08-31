package database

import "yalk/database/models"

type DatabaseOperations interface {
	SaveMessage(*models.Message) error
	GetUsers(string) ([]string, error)
}
