package usecase

import "asgo/interfaces/database"

type DataInteractor struct {
	DataRepo DataRepository
}

func (interactor *DataInteractor) Create(data *database.Data) error {
	if err := interactor.DataRepo.InsertData(data); err != nil {
		return err
	}
	return nil
}

func (interactor *DataInteractor) Select(userid string) (*database.Data, error) {
	data, err := interactor.DataRepo.SelectDataByUserID(userid)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (interactor *DataInteractor) Update(data *database.Data) error {
	if err := interactor.DataRepo.UpdateData(data); err != nil {
		return err
	}
	return nil
}
