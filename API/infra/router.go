package infra

import (
	"asgo/infra/action"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"*",
		},
	}))

	asgoreRoute := router.Group("/asgore")
	{
		asgoreRoute.GET("/data", func(c *gin.Context) {

		})

		asgoreRoute.POST("/signup", func(c *gin.Context) {

		})
	}

	bonusRoute := router.Group("/bonus")
	{
		bonusRoute.GET("/number", func(c *gin.Context) {
			action.GetBonusNumber(c)
		})
		bonusRoute.GET("/week", func(c *gin.Context) {
			action.GetBonusWeek(c)
		})
		bonusRoute.GET("/roll", func(c *gin.Context) {
			action.GetBonusRoll(c)
		})
	}

	ticketRoute := router.Group("/ticket")
	{
		ticketRoute.GET("/number", func(c *gin.Context) {
			action.GetTicketNumber(c)
		})
		ticketRoute.GET("/roll", func(c *gin.Context) {
			action.GetTicketRoll(c)
		})
	}

	questionnaireRoute := router.Group("/questionnaire")
	{
		questionnaireRoute.GET("/content", func(c *gin.Context) {
			action.GetContent(c)
		})
		questionnaireRoute.POST("/autowrite", func(c *gin.Context) {
			action.PostWrite(c)
		})
	}
}
