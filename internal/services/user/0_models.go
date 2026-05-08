package user

import "context"

type LoginRequest struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpiry  int64  `json:"refresh_token_expiry"`
	IsAdmin             bool   `json:"is_admin"`
	Email               string `json:"email"`
}

type LoginServiceParams struct {
	ctx context.Context
	GoogleIDToken string
	IP string
	UserAgent string
}