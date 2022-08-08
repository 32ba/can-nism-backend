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
		}
		tokenEngine := apiV1.Group("/token")
		{
			tokenEngine.POST("/refresh", controller.RefreshAccessToken)
		}
		rankingEngine := apiV1.Group("/ranking")
		{
			rankingEngine.POST("/add", controller.AddRanking)
		}
	}

	engine.Run(":8080")
}
