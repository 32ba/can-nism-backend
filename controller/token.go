package controller

import (
	"go-ranking-api/ent/token"
	"go-ranking-api/model"
	"go-ranking-api/repository"
	"go-ranking-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RefreshAccessToken(c *gin.Context) {
	var accessTokenRefreshRequest repository.AccessTokenRefreshRequest
	if err := c.ShouldBindJSON(&accessTokenRefreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if accessTokenRefreshRequest.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token not found"})
		return
	}

	refreshToken, err := model.DBClient.Token.
		Query().
		Where(
			token.RefreshTokenEQ(accessTokenRefreshRequest.RefreshToken),
			token.RefreshTokenExpiredAtGTE(time.Now()),
			token.HasUser(),
		).
		Only(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := utils.RandomStringGenerator(16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := model.DBClient.Token.
		UpdateOneID(refreshToken.ID).
		SetAccessToken(accessToken).
		SetAccessTokenExpiredAt(time.Now().Add(1 * time.Hour)).
		Save(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res repository.AccessTokenRefreshResponse
	res.AccessToken = token.AccessToken
	c.JSON(http.StatusOK, res)
}