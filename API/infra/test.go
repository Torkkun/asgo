package infra

import (
	"asgo/usecase"
)

// 新しくrepositoryをつくる
func Test() {
	selehandler := NewSeleHandler()
	sqlhandler := NewSqlHandler()
	repository := usecase.NewRepository(sqlhandler, selehandler)
	_ = testNewUserController(repository)

}

type UserController struct {
	SInteracter usecase.SecretInteractor
	UInteractor usecase.UserInteractor
}

func testNewUserController(repository *usecase.Repositorys) *UserController {
	return &UserController{
		SInteracter: usecase.SecretInteractor{
			SecretRepo: repository.SecretRepo,
		},
		UInteractor: usecase.UserInteractor{
			UserRepo: repository.UserRepo,
		},
	}
}
