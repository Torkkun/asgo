package usecase

import "asgo/domain"

type UserInteractor struct {
	UserRepo UserRepository
}

func (interactor *UserInteractor) Create(u domain.UserDB) (err error) {
	_, err = interactor.UserRepo.InsertUser(u)
	return
}

func (interactor *UserInteractor) UserById(identifier string) (user domain.UserDB, err error) {
	user, err = interactor.UserRepo.SelectUserFindById(identifier)
	return
}
