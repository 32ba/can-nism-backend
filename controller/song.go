package controller

import (
	"go-ranking-api/ent/song"
	"go-ranking-api/model"
	"go-ranking-api/repository"
	"go-ranking-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSongList(c *gin.Context) {
	tokenID, err := utils.TokenVerifier("access", c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tokenID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token velification error"})
		return
	}

	songs, err := model.DBClient.Song.
		Query().
		Where(
			song.DeletedAtIsNil(),
		).
		All(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var r []repository.Song
	for _, e := range songs {
		var elem repository.Song
		elem.UUID = e.UUID
		elem.Title = e.Title
		r = append(r, elem)
	}
	var res repository.GetSongsResponse
	res.Songs = r

	c.JSON(http.StatusOK, res)
}
