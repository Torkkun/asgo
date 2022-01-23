package infra

import (
	"asgo/interfaces/controllers"

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

	selehandler := NewSeleHandler()
	seleController := controllers.NewSeleController(selehandler)

	asgoreRoute := router.Group("/asgore")
	{

		asgoreRoute.POST("/signup", func(c *gin.Context) {

		})

		userRoute := asgoreRoute.Group("/user")
		{
			userRoute.GET("/new", func(c *gin.Context) {

			})
			userRoute.POST("/create", func(c *gin.Context) {

			})
		}

		scrapingRoute := asgoreRoute.Group("/scrap")
		{
			scrapingRoute.GET("/data", func(c *gin.Context) {
				seleController.Scraping(c)
			})
			scrapingRoute.GET("/daily", func(c *gin.Context) {
				seleController.Roll(c)
			})
			scrapingRoute.GET("/ticket", func(c *gin.Context) {

			})
			scrapingRoute.GET("/bonus", func(c *gin.Context) {

			})

		}

		dataRoute := asgoreRoute.Group("/data")
		{
			dataRoute.POST("/create", func(c *gin.Context) {

			})
			dataRoute.GET("/read", func(c *gin.Context) {

			})

			dataRoute.PUT("/update", func(c *gin.Context) {

			})

			dataRoute.DELETE("/delete", func(c *gin.Context) {

			})
		}
	}

	questionnaireRoute := router.Group("/questionnaire")
	{
		questionnaireRoute.GET("/content", func(c *gin.Context) {
		})
		questionnaireRoute.POST("/autowrite", func(c *gin.Context) {
		})
	}
}
