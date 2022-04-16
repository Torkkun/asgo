package controllers

import (
	"asgo/domain"
	"asgo/interfaces/database"
	"asgo/usecase"
	"log"
)

type UserController struct {
	DBService usecase.DBInteractor
}

func NewUserController(sqlhandler database.SqlHandler) *UserController {
	dbService := usecase.NewDBService(sqlhandler)
	return &UserController{
		DBService: *dbService}
}

// discordのuidを取得しワンタイムパスワードを作成　// 時間制限をどうするか
func (con *UserController) NewCreate(c Context) {
	identifier := c.Query("client_uid")
	oneTimePass, err := createpasscode()
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
	}
	if err = con.DBService.CreateSecret(&database.Secret{ClientUserID: identifier, OneTimePassWord: oneTimePass}); err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
	}
	c.JSON(201, domain.SecretResponse{Code: oneTimePass})
}

func (controller *UserController) Create(c Context) {

}

func (controller *UserController) Index(c Context) {

}

func (controller *UserController) Show(c Context) {

}
