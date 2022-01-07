package usecase

import "asgo/domain"

type ActionInteractor struct {
	ActionRepository ActionRepository
}

func (interactor *ActionInteractor) DailyGatya(count int) (daily domain.DailyGatya, err error) {
	daily, err = interactor.ActionRepository.DailyRoll(count)
	return
}
