package usecase

import "asgo/domain"

//seleniumでスクレイピングしてくる
type SeleRepository interface {
	Login(domain.User) error
	DailyRoll() (domain.DailyGatya, error)
	BonusRoll(int) error
	ExchangeRoll(int) error
	CheckTicket() error
}
