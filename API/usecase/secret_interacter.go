package usecase

import "asgo/interfaces/database"

func (interactor *DBInteractor) CreateSecret(user *database.Secret) error {
	if err := interactor.SecretRepo.InsertSecret(user); err != nil {
		return err
	}
	return nil
}
