package controller

import (
	"fmt"
	"go-ranking-api/ent"
	"go-ranking-api/ent/ranking"
	"go-ranking-api/ent/user"
	"go-ranking-api/model"
	"go-ranking-api/repository"
	"go-ranking-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//var String[] supportedRankingTypes = [2]string {"top5", "top100"}

func rankingTypes(cat string) int {
	switch cat {
	case "top5":
		return 5
	case "top100":
		return 100
	default:
		return 0
	}
}

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

	u, err := utils.GetUserFromTokenID(tokenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isFoundRecord, err := model.DBClient.Ranking.
		Query().
		Where(
			ranking.And(
				ranking.SongUUIDEQ(addRankingRequest.SongUUID),
				ranking.HasUserWith(user.ID(u.ID)),
			),
		).
		Exist(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isFoundRecord {
		err := model.DBClient.Ranking.
			Update().
			Where(
				ranking.And(
					ranking.SongUUIDEQ(addRankingRequest.SongUUID),
					ranking.HasUserWith(user.ID(u.ID)),
				),
			).
			SetScore(addRankingRequest.Score).
			Exec(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}else{
		err := model.DBClient.Ranking.
			Create().
			SetScore(addRankingRequest.Score).
			SetSongUUID(addRankingRequest.SongUUID).
			SetUser(u).
			Exec(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func GetRanking(c *gin.Context) {
	var req repository.GetRankingRequest
	if c.Query("song_uuid") == "" || c.Query("ranking_type") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query strings invalid"})
		return
	}
	if rankingTypes(c.Query("ranking_type")) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ranking type"})
		return
	}
	req.Type = c.Query("ranking_type")
	req.SongUUID = uuid.MustParse(c.Query("song_uuid"))

	ranking, err := model.DBClient.Ranking.
		Query().
		Where(
			ranking.SongUUIDEQ(req.SongUUID),
		).
		WithUser().
		Order(ent.Desc(ranking.FieldScore)).
		Limit(rankingTypes(req.Type)).
		All(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(ranking)
	var r []repository.Ranking
	for i, e := range ranking {
		var elem repository.Ranking
		elem.Name = e.Edges.User.Name
		elem.Rank = i + 1
		elem.Score = int(e.Score)
		r = append(r, elem)
	}

	var res repository.GetRankingResponse
	res.SongUUID = req.SongUUID
	res.RankingType = req.Type
	res.Ranking = r

	c.JSON(http.StatusOK, res)
}
