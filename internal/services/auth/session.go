package auth

import "github.com/google/uuid"

type AuthResponse struct {
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpiry  int64  `json:"refresh_token_expiry"`
	IsAdmin             bool   `json:"is_admin"`
	Email               string `json:"email"`
}

type AuthSession struct {
	SessionID uuid.UUID
	UserID uuid.UUID
}