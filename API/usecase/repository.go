package usecase

import (
	"asgo/interfaces/database"
	"asgo/interfaces/selenium"
)

type Repositorys struct {
	SecretRepo     SecretRepository
	UserRepo       UserRepository
	SeleRepository SeleRepository
}

//まとめてrepositoryを返す　どちらが良いかは後にしらべる
func NewRepository(sqlHandler database.SqlHandler, selehandler selenium.SeleHandler) *Repositorys {
	return &Repositorys{
		SecretRepo: &database.SecretRepository{
			SqlHandler: sqlHandler,
		},
		UserRepo: &database.UserRepository{
			SqlHandler: sqlHandler,
		},
		SeleRepository: &selenium.SeleniumRepository{
			SeleHandler: selehandler,
		},
	}
}
