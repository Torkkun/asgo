package usecase

import "asgo/domain"

type SecretRepository interface {
	InsertSecret(domain.SecretCode) (int, error)
}
