package usecase

import "asgo/interfaces/database"

type SecretInteractor struct {
	SecretRepo SecretRepository
}

func (interactor *SecretInteractor) Create(user *database.Secret) error {
	if err := interactor.SecretRepo.InsertSecret(user); err != nil {
		return err
	}
	return nil
}
