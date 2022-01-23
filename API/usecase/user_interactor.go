package usecase

import "asgo/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) NewCreate(u domain.SecretCode) (err error) {
	_, err = interactor.UserRepository.StoreCode(u)
	return
}

func (interactor *UserInteractor) Create(u domain.UserDB) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) UserById(identifier string) (user domain.UserDB, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
