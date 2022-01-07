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

	asgoreRoute := router.Group("/asgore")
	{
		asgoreRoute.GET("/data", func(c *gin.Context) {

		})

		asgoreRoute.POST("/signup", func(c *gin.Context) {

		})

		asgoreRoute.GET("/daily", func(c *gin.Context) {

		})

		asgoreRoute.GET("/ticket", func(c *gin.Context) {

		})

		asgoreRoute.GET("/bonus", func(c *gin.Context) {

		})
	}

	questionnaireRoute := router.Group("/questionnaire")
	{
		questionnaireRoute.GET("/content", func(c *gin.Context) {
		})
		questionnaireRoute.POST("/autowrite", func(c *gin.Context) {
		})
	}
}
