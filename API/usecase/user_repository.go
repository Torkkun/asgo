package usecase

import "asgo/domain"

type UserRepository interface {
	Store(domain.UserDB) (int, error)
	StoreCode(domain.SecretCode) (int, error)
	FindById(string) (domain.UserDB, error)
	FindCode(string) (domain.SecretCode, error)
}
