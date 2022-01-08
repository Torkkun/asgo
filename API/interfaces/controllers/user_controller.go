package controllers

import (
	"asgo/interfaces/database"
	"asgo/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {

}

func (controller *UserController) Index(c Context) {

}

func (controller *UserController) Show(c Context) {

}
