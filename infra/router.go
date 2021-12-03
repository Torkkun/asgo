package infra

import (
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
	//そのうち抽象化する
	gatyaRoute := router.Group("/gatya")
	{
		gatyaRoute.GET("/point", func(c *gin.Context) {
			getpoint(c)
		})
		gatyaRoute.GET("/roll", func(c *gin.Context) {
			getroll(c)
		})
		gatyaRoute.GET("rolls", func(c *gin.Context) {
			getrolls(c)
		})
	}

	bonusRoute := router.Group("/bonus")
	{
		bonusRoute.GET("/number", func(c *gin.Context) {
			getbonusnumber(c)
		})
		bonusRoute.GET("/week", func(c *gin.Context) {
			getbonusweek(c)
		})
		bonusRoute.GET("/roll", func(c *gin.Context) {
			getbonusroll(c)
		})
		bonusRoute.GET("rolls", func(c *gin.Context) {
			getbonusrolls(c)
		})
	}

	ticketRoute := router.Group("/ticket")
	{
		ticketRoute.GET("/number", func(c *gin.Context) {
			getticketnumber(c)
		})
		ticketRoute.GET("/roll", func(c *gin.Context) {
			getticketroll(c)
		})
		ticketRoute.GET("/rolls", func(c *gin.Context) {
			getticketrolls(c)
		})
	}

	questionnaireRoute := router.Group("/questionnaire")
	{
		questionnaireRoute.GET("/content", func(c *gin.Context) {
			getcontent(c)
		})
		questionnaireRoute.POST("/autowrite", func(c *gin.Context) {
			postwrite(c)
		})
	}
}
