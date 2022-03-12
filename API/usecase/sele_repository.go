package usecase

import (
	"asgo/interfaces/selenium"
)

//seleniumでスクレイピングしてくる
type SeleRepository interface {
	Login(*selenium.Login) error
	DailyRoll() (*selenium.DailyGatya, error)
	BonusRoll(int) error
	ExchangeRoll(int) error
	CheckTicket() error
}
