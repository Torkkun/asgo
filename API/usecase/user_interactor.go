package usecase

import (
	"asgo/interfaces/database"
	"log"
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

func (interactor *UserInteractor) UserById(identifier string) (*database.User, error) {
	user, err := interactor.UserRepo.SelectUserFindById(identifier)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
