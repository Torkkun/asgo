package usecase

import "asgo/interfaces/database"

type SecretRepository interface {
	InsertSecret(*database.Secret) error
	SelectSecretById(string) (*database.Secret, error)
}
