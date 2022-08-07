package repository

type UserSignupResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken string    `json:"access_token"`
}

type AccessTokenRefreshResponse struct {
	AccessToken string `json:"access_token"`
}