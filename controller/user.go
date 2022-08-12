package controller

import (
	"go-ranking-api/model"
	"go-ranking-api/repository"
	"go-ranking-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UserSignup(c *gin.Context) {
	var signupRequest repository.UserSignupRequest
	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := utils.RandomStringGenerator(16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err := utils.RandomStringGenerator(64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.DBClient.User.
		Create().
		SetGUID(signupRequest.GUID).
		SetName(signupRequest.Name).
		Save(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := model.DBClient.Token.
		Create().
		SetUser(user).
		SetAccessToken(accessToken).
		SetAccessTokenExpiredAt(time.Now().Add(12 * time.Hour)).
		SetRefreshToken(refreshToken).
		SetRefreshTokenExpiredAt(time.Now().AddDate(1, 0, 0)).
		Save(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response repository.UserSignupResponse
	response.AccessToken = token.AccessToken
	response.RefreshToken = token.RefreshToken

	c.JSON(http.StatusOK, response)
}

func UpdateUserName(c *gin.Context) {
	var updateUserNameRequest repository.UpdateUserNameRequest
	if err := c.ShouldBindJSON(&updateUserNameRequest); err != nil {
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

	u, err := utils.GetUserFromTokenID(tokenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = model.DBClient.User.
		UpdateOneID(u.ID).
		SetName(updateUserNameRequest.Name).
		Save(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
