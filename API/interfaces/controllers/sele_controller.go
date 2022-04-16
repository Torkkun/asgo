package controllers

import (
	"asgo/domain"
	"asgo/interfaces/database"
	"asgo/interfaces/selenium"
	"asgo/usecase"
	"time"

	"log"
)

type SakitoController struct {
	SakitoService usecase.SeleniumInteractor
	DBService     usecase.DBInteractor
}

func NewSakitoController(selehandler selenium.SeleHandler, sqlhandler database.SqlHandler) *SakitoController {
	seleService := usecase.NewSeleService(selehandler)
	dbService := usecase.NewDBService(sqlhandler)
	return &SakitoController{
		SakitoService: *seleService,
		DBService:     *dbService,
	}
}

func (controller *SakitoController) Roll(c Context) {
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
	udata, err := controller.DBService.SelectUserById(userid)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	// DBからデータを取得し整形
	data, err := controller.SakitoService.DailyGatya(
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
func (controller *SakitoController) HomeData(c Context) {
	userid := GetUserIDFromContext(c)
	if userid == "" {
		log.Println("unauthorized")
		c.JSON(401, "Error: Unauthorized")
		return
	}
	// DBからデータを取得
	data, err := controller.DBService.GetData(userid)
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
	udata, err := controller.DBService.SelectUserById(userid)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	// dataをスクレイピング
	mydata, err := controller.SakitoService.MyData(&selenium.Login{
		Email:    udata.Email,
		Password: udata.Password})
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
		return
	}
	if data == nil {
		// データが存在しないとき
		if err := controller.DBService.CreateData(
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
		if err := controller.DBService.UpdateData(
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
