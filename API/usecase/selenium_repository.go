package usecase

import "asgo/domain"

//seleniumでスクレイピングしてくる
type ActionRepository interface {
	DailyRoll(int) (domain.DailyGatya, error)
	BonusRoll(int)
	ExchangeRoll(int)
	CheckTicket()
}
