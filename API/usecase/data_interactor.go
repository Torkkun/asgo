package usecase

import (
	"asgo/interfaces/database"
)

func (interactor *DBInteractor) CreateData(data *database.Data) error {
	if err := interactor.DataRepo.InsertData(data); err != nil {
		return err
	}
	return nil
}

func (interactor *DBInteractor) GetData(userid string) (*database.Data, error) {
	data, err := interactor.DataRepo.SelectDataByUserID(userid)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (interactor *DBInteractor) UpdateData(data *database.Data) error {
	if err := interactor.DataRepo.UpdateData(data); err != nil {
		return err
	}
	return nil
}
