package usecase

import (
	"asgo/interfaces/database"
)

type UserRepository interface {
	InsertUser(*database.User) error
	SelectUserFindById(string) (*database.User, error)
}
