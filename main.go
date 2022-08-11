package main

import (
	"go-ranking-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/healthchecker", controller.HealthChecker)
	apiV1 := engine.Group("/api/v1")
	{
		userEngine := apiV1.Group("/user")
		{
			userEngine.POST("/signup", controller.UserSignup)
			userEngine.POST("/update")
		}
		tokenEngine := apiV1.Group("/token")
		{
			tokenEngine.POST("/refresh", controller.RefreshAccessToken)
		}
		rankingEngine := apiV1.Group("/ranking")
		{
			rankingEngine.POST("", controller.AddRanking)
			rankingEngine.GET("", controller.GetRanking)
		}
		songEngine := apiV1.Group("/song")
		{
			songEngine.GET("/list", controller.GetSongList)
		}
	}

	engine.Run(":8080")
}
