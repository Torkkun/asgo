package usecase

import (
	"asgo/interfaces/database"
	"asgo/interfaces/selenium"
)

type DBInteractor struct {
	SecretRepo SecretRepository
	UserRepo   UserRepository
	DataRepo   DataRepository
}

type SeleniumInteractor struct {
	SeleRepo SeleRepository
}

// repositoryをusecaseintaractorで使用できるようにする
//まとめてrepositoryを返す　サービスとしてインジェクションする
func NewDBService(sqlHandler database.SqlHandler) *DBInteractor {
	return &DBInteractor{
		SecretRepo: &database.SecretRepository{
			SqlHandler: sqlHandler,
		},
		UserRepo: &database.UserRepository{
			SqlHandler: sqlHandler,
		},
		DataRepo: &database.DataRepository{
			SqlHandler: sqlHandler,
		},
	}
}

func NewSeleService(selehandler selenium.SeleHandler) *SeleniumInteractor {
	return &SeleniumInteractor{
		SeleRepo: &selenium.SeleniumRepository{
			SeleHandler: selehandler,
		},
	}
}
