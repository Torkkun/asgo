package controllers

import (
	"asgo/domain"
	"asgo/interfaces/selenium"
	"asgo/usecase"

	"log"
)

type SeleController struct {
	Interactor usecase.SeleInteractor
}

func NewSeleController(selehandler selenium.SeleHandler) *SeleController {
	return &SeleController{
		Interactor: usecase.SeleInteractor{
			SeleRepository: &selenium.SeleniumRepository{
				SeleHandler: selehandler,
			},
		},
	}
}

func (controller *SeleController) Roll(c Context) {
	//userid := GetUserIDFromContext(c)
	//useridを使用してdatabaseからデータを取得して
	user := domain.UsersRequest{}
	err := c.Bind(&user)
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
		return
	}
	// DBからデータを取得し整形
	datas, err := controller.Interactor.DailyGatya(&usecase.SakitoLogin{})
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, datas)
}

func (controller *SeleController) Scraping(c Context) {
	//userid := c.Query("userid")

}
