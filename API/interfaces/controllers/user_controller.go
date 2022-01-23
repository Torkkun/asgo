package controllers

import (
	"asgo/domain"
	"asgo/interfaces/database"
	"asgo/usecase"
	"log"
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

func (controller *UserController) NewCreate(c Context) {
	identifier := c.Query("client_uid")
	onetimepass, err := createpasscode()
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
	}
	userdata := domain.SecretCode{
		ClientUserID:    identifier,
		OneTimePassWord: onetimepass,
	}
	err = controller.Interactor.NewCreate(userdata)
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
	}
	code := domain.Code{
		Code: onetimepass,
	}
	c.JSON(201, code)
}

func (controller *UserController) Create(c Context) {

}

func (controller *UserController) Index(c Context) {

}

func (controller *UserController) Show(c Context) {

}
