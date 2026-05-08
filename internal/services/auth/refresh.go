package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"go-tailwind-test/internal/config"
	"time"

	"github.com/google/uuid"
)

// only for testing

func GenerateRefreshToken() string {
	return uuid.NewString()
}

func HashRefreshToken(token string,) string {
	hash := sha256.Sum256(
		[]byte(token),
	)

	return hex.EncodeToString(hash[:])
}

func GenerateRefreshTokenExpiry() int64 {
	var RefreshTokenDuration = time.Duration(config.Envs.JWTRefreshExpirationMinutes) * time.Minute
	return time.Now().Add(RefreshTokenDuration).Unix()
}

func IsRefreshTokenExpired(expiresAt int64,) bool {
	return time.Now().Unix() > expiresAt
}