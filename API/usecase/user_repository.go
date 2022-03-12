package usecase

import "asgo/domain"

type UserRepository interface {
	InsertUser(domain.UserDB) (int, error)
	SelectUserFindById(string) (domain.UserDB, error)
	SelectUserFindByCode(string) (domain.SecretCode, error)
}
