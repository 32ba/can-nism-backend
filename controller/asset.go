package controller

import (
	"fmt"
	"go-ranking-api/ent/asset"
	"go-ranking-api/ent/song"
	"go-ranking-api/model"
	"go-ranking-api/repository"
	"go-ranking-api/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func platformTypes(p string) asset.Platform {
	switch p {
	case "windows":
		return asset.PlatformWINDOWS
	case "ios":
		return asset.PlatformIOS
	default:
		return asset.Platform("NULL")
	}
}

func GetAssetURL(c *gin.Context) {
	tokenID, err := utils.TokenVerifier("access", c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tokenID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token velification error"})
		return
	}

	var getAssetURLRequest repository.GetAssetUrlRequest
	if err := c.ShouldBindJSON(&getAssetURLRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if platformTypes(getAssetURLRequest.Platform) == asset.Platform("NULL") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid platfrom"})
		return
	}

	asset, err := model.DBClient.Asset.
		Query().
		Where(
			asset.PlatformEQ(platformTypes(getAssetURLRequest.Platform)),
			asset.HasSongWith(song.UUID(getAssetURLRequest.SongUUID)),
			asset.DeletedAtIsNil(),
		).
		First(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res repository.GetAssetUrlResponse
	res.URL = fmt.Sprintf("%s/%s/%s", os.Getenv("STATIC_BASE_URL"), getAssetURLRequest.Platform, asset.Path)
	res.Hash = asset.Hash

	c.JSON(http.StatusOK, res)
}
