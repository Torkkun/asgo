package controllers

import (
	"asgo/domain"
	"asgo/interfaces/database"
	"asgo/interfaces/selenium"
	"asgo/usecase"
	"time"

	"log"
)

type SeleController struct {
	SeleInteractor usecase.SeleInteractor
	UInteractor    usecase.UserInteractor
	DInteractor    usecase.DataInteractor
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
		&selenium.Login{
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

// データを取得
func (controller *SeleController) ScrapingData(c Context) {
	userid := GetUserIDFromContext(c)
	if userid == "" {
		log.Println("unauthorized")
		c.JSON(401, "Error: Unauthorized")
		return
	}
	// DBからデータを取得
	data, err := controller.DInteractor.Select(userid)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	// データーが存在しており更新日時が一分以内の場合
	if data != nil && !checkTime(data.UpdatedAt) {
		c.JSON(200, domain.DataResponse{
			DayPoint:    data.DailyPoint,
			Points:      data.Points,
			BonusTicket: data.BonusTicket,
			BonusWeek:   data.BonusWeek,
		})
		return
	}
	// userのpassとemailを取得
	udata, err := controller.UInteractor.SelectUserById(userid)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	// dataをスクレイピング
	mydata, err := controller.SeleInteractor.MyData(&selenium.Login{
		Email:    udata.Email,
		Password: udata.Password})
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	if data == nil {
		// データが存在しないとき
		if err := controller.DInteractor.Create(
			&database.Data{
				UserID:      userid,
				Points:      mydata.Points,
				BonusTicket: mydata.BonusTicket,
				BonusWeek:   mydata.BonusWeek,
			}); err != nil {
			c.JSON(500, "Internal Server Error")
			return
		}
		c.JSON(201, domain.DataResponse{
			DayPoint:    0,
			Points:      mydata.Points,
			BonusTicket: mydata.BonusTicket,
			BonusWeek:   mydata.BonusWeek,
		})
	} else {
		// データが存在するときアップデート
		if err := controller.DInteractor.Update(
			&database.Data{
				Points:      mydata.Points,
				BonusTicket: mydata.BonusTicket,
				BonusWeek:   mydata.BonusWeek,
			}); err != nil {
			c.JSON(500, "Internal Server Error")
			return
		}
		c.JSON(200, domain.DataResponse{
			DayPoint:    data.DailyPoint,
			Points:      mydata.Points,
			BonusTicket: mydata.BonusTicket,
			BonusWeek:   mydata.BonusWeek,
		})
	}
}

func checkTime(t time.Time) bool {
	nowduration := time.Since(t)
	duration, err := time.ParseDuration("1m")
	if err != nil {
		log.Fatalln(err)
	}
	return nowduration >= duration
}
