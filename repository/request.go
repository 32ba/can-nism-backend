package repository

import "github.com/google/uuid"

type UserSignupRequest struct {
	GUID uuid.UUID `json:"guid"`
	Name string    `json:"name"`
}

type AccessTokenRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AddRankingRequest struct {
	Score int64 `json:"score"`
}
