package usecase

import (
	"asgo/interfaces/database"
)

func (interactor *DBInteractor) CreateUser(user *database.User) error {
	if err := interactor.UserRepo.InsertUser(user); err != nil {
		return err
	}
	return nil
}

func (interactor *DBInteractor) SelectUserById(userID string) (*database.User, error) {
	user, err := interactor.UserRepo.SelectUserFindById(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
