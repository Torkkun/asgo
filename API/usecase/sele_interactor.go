package usecase

import (
	"asgo/interfaces/selenium"
)

type SakitoLogin struct {
	Email    string
	Password string
}

func (interactor *SeleniumInteractor) DailyGatya(slogin *selenium.Login) (*selenium.DailyGatya, error) {
	if err := interactor.SeleRepo.Login(
		&selenium.Login{
			Email:    slogin.Email,
			Password: slogin.Password}); err != nil {
		return nil, err
	}
	daily, err := interactor.SeleRepo.DailyRoll()
	if err != nil {
		return nil, err
	}
	return daily, err
}

func (interactor *SeleniumInteractor) MyData(slogin *selenium.Login) (*selenium.Data, error) {
	var data selenium.Data
	// login
	if err := interactor.SeleRepo.Login(&selenium.Login{
		Email:    slogin.Email,
		Password: slogin.Password,
	}); err != nil {
		return nil, err
	}
	// dataをスクレイピング
	return &data, nil
}
