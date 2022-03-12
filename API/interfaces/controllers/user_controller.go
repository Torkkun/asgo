package controllers

import (
	"asgo/domain"
	"asgo/interfaces/database"
	"asgo/usecase"
	"log"
)

type UserController struct {
	SInteracter usecase.SecretInteractor
	UInteractor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		SInteracter: usecase.SecretInteractor{
			SecretRepo: &database.SecretRepository{
				SqlHandler: sqlHandler,
			},
		},
		UInteractor: usecase.UserInteractor{
			UserRepo: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// discordのuidを取得しワンタイムパスワードを作成　// 時間制限をどうするか
func (con *UserController) NewCreate(c Context) {
	identifier := c.Query("client_uid")
	oneTimePass, err := createpasscode()
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
	}

	if err = con.SInteracter.Create(&database.Secret{ClientUserID: identifier, OneTimePassWord: oneTimePass}); err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
	}
	c.JSON(201, domain.SecretResponse{Code: oneTimePass})
}

func (controller *UserController) Create(c Context) {

}

func (controller *UserController) Index(c Context) {

}

func (controller *UserController) Show(c Context) {

}
