package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserEmail string `json:"user_email"`
	IsAdmin   bool   `json:"is_admin"`
	SessionID  uuid.UUID `json:"session_id"`
	jwt.RegisteredClaims
}

