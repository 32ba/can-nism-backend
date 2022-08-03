package main

import (
	"go-ranking-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/healthchecker", controller.HealthChecker)

	engine.Run(":8080")
}