package repository

import "github.com/google/uuid"

type UserSignupResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type AccessTokenRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type Ranking struct {
	Rank  int    `json:"rank"`
	Score int    `json:"score"`
	Name  string `json:"name"`
}
type GetRankingResponse struct {
	SongUUID    uuid.UUID `json:"song_uuid"`
	RankingType string    `json:"ranking_type"`
	Ranking     []Ranking `json:"ranking"`
}

type Song struct {
	UUID  uuid.UUID `json:"uuid"`
	Title string    `json:"name"`
}
type GetSongsResponse struct {
	Songs []Song `json:"songs"`
}

type GetAssetUrlResponse struct {
	URL  string `json:"url"`
	Hash string `json:"hash"`
}
