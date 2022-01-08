package usecase

import (
	"asgo/domain"
	"log"
)

type SeleInteractor struct {
	SeleRepository SeleRepository
}

func (interactor *SeleInteractor) DailyGatya(users []domain.User) (datas domain.DailyDatas, err error) {
	if len(users) > 0 {
		for _, user := range users {
			err = interactor.SeleRepository.Login(user)
			if err != nil {
				return
			}
			daily, err := interactor.SeleRepository.DailyRoll()
			if err != nil {
				log.Println(err)
				continue
			}
			datas = append(datas, daily)
		}
	}
	return
}

func (interactor *SeleInteractor) MyData(userid string) (err error) {
	return
}
