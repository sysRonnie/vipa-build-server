package auth

import (
	"errors"
	"go-tailwind-test/internal/config"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateAccessToken(
	email string,
	isAdmin bool,
	sessionID uuid.UUID,
) (string, error) {

	now := time.Now()

	cfg := config.Envs


	claims := Claims{
		UserEmail: email,
		IsAdmin:   isAdmin,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(
				now.Add(time.Duration(cfg.JWTExpirationMinutes) * time.Minute),
			),
		},
	}


	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	log.Println("[GenerateAccessToken] jwt token = ", token)
	log.Println("[GenerateAccessToken] jwt expiration time = ",jwt.NewNumericDate(
				now.Add(time.Duration(cfg.JWTExpirationMinutes) * time.Minute) ))

	signedToken, err := token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseAccessToken(
	tokenString string,
) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(
					"invalid signing method",
				)
			}

			return []byte(
				os.Getenv("JWT_SECRET"),
			), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, errors.New(
			"invalid token claims",
		)
	}

	if !token.Valid {
		return nil, errors.New(
			"invalid token",
		)
	}

	return claims, nil
}