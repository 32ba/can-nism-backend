package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthChecker(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
