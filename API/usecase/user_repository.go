package usecase

import "asgo/domain"

type UserRepository interface {
	Store(domain.UserDB) (int, error)
	FindById(int) (domain.UserDB, error)
	FindAll() (domain.UsersDB, error)
}
