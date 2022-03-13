package usecase

import "asgo/interfaces/database"

type ClientRepository interface {
	InsertClient(*database.Client)
	SelectClientFindByID(string) (*database.Client, error)
}
