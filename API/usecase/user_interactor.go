package usecase

import "asgo/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.UserDB) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Users() (user domain.UsersDB, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.UserDB, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
