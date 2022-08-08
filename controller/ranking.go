package controller

import (
	"go-ranking-api/model"
	"go-ranking-api/repository"
	"go-ranking-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRanking(c *gin.Context) {
	var addRankingRequest repository.AddRankingRequest
	if err := c.ShouldBindJSON(&addRankingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenID, err := utils.TokenVerifier("access", c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tokenID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token velification error"})
		return
	}

	user, err := utils.GetUserFromTokenID(tokenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = model.DBClient.Ranking.Create().
		SetScore(addRankingRequest.Score).
		SetUser(user).
		OnConflict().
		UpdateScore().
		Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
