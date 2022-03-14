package usecase

import "asgo/interfaces/database"

type DataRepository interface {
	InsertData(*database.Data) error
	SelectDataByUserID(string) (*database.Data, error)
	UpdateData(*database.Data) error
}
