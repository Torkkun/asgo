package controllers

import (
	"asgo/domain"
	"asgo/interfaces/database"
	"asgo/interfaces/selenium"
	"asgo/usecase"

	"log"
)

type SeleController struct {
	SeleInteractor usecase.SeleInteractor
	UInteractor    usecase.UserInteractor
}

func NewSeleController(selehandler selenium.SeleHandler, sqlhandler database.SqlHandler) *SeleController {
	return &SeleController{
		SeleInteractor: usecase.SeleInteractor{
			SeleRepo: &selenium.SeleniumRepository{
				SeleHandler: selehandler,
			},
		},
		UInteractor: usecase.UserInteractor{
			UserRepo: &database.UserRepository{
				SqlHandler: sqlhandler,
			},
		},
	}
}

func (controller *SeleController) Roll(c Context) {
	userid := GetUserIDFromContext(c)
	if userid == "" {
		log.Println("unauthorized")
		c.JSON(401, "Error: Unauthorized")
		return
	}
	user := domain.UsersRequest{}
	err := c.Bind(&user)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	// userのpassとemailを取得
	udata, err := controller.UInteractor.SelectUserById(userid)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	// DBからデータを取得し整形
	data, err := controller.SeleInteractor.DailyGatya(
		&usecase.SakitoLogin{
			Email:    udata.Email,
			Password: udata.Password})
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(201, domain.DailyGatyaResponse{
		DayPoint:      data.DayPoint,
		PointSum:      data.PointSum,
		ExecutionDate: data.ExecutionDate,
	})
}

func (controller *SeleController) ScrapingData(c Context) {

}
