package usecase

import (
	"asgo/interfaces/database"
)

type UserInteractor struct {
	UserRepo UserRepository
}

func (interactor *UserInteractor) Create(user *database.User) error {
	if err := interactor.UserRepo.InsertUser(user); err != nil {
		return err
	}
	return nil
}

func (interactor *UserInteractor) SelectUserById(userID string) (*database.User, error) {
	user, err := interactor.UserRepo.SelectUserFindById(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
