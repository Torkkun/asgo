package usecase

import "asgo/domain"

type SecretInteractor struct {
	SecretRepo SecretRepository
}

func (interactor *SecretInteractor) Create(u domain.SecretCode) (err error) {
	_, err = interactor.SecretRepo.InsertSecret(u)
	return
}
