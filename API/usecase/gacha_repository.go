package usecase

import "asgo/domain"

type GatyaRepository interface {
	FindById() (domain.Gatya, error)
}
