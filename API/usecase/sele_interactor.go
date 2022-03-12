package usecase

import (
	"asgo/interfaces/selenium"
)

type SeleInteractor struct {
	SeleRepository SeleRepository
}

type SakitoLogin struct {
	Email    string
	Password string
}

func (interactor *SeleInteractor) DailyGatya(slogin *SakitoLogin) (*selenium.DailyGatya, error) {
	if err := interactor.SeleRepository.Login(&selenium.Login{Email: slogin.Email, Password: slogin.Password}); err != nil {
		return nil, err
	}
	daily, err := interactor.SeleRepository.DailyRoll()
	if err != nil {
		return nil, err
	}
	return daily, err
}

func (interactor *SeleInteractor) MyData(userid string) (err error) {
	return
}
