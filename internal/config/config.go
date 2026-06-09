package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string
	DatabaseDSN string
	GoogleClientIDWeb string
	GoogleClientIDAndroid string
	JWTExpirationMinutes int
	JWTRefreshExpirationMinutes int
}

var Envs = InitializeConfig()

func InitializeConfig() Config {
	_ = godotenv.Load()

	env := strings.ToLower(getEnv("APP_ENV", "dev"))
	databaseDSN := mustGetEnv("DATABASE_DSN")
	googleClientIDWeb := mustGetEnv("GOOGLE_CLIENT_ID_WEB")
	googleClientIDAndroid := mustGetEnv("GOOGLE_CLIENT_ID_ANDROID")
	jwtExpirationMinutes := mustGetInt("JWT_EXPIRATION_MINUTES")
	jwtRefreshExpirationDays := mustGetInt("REFRESH_TOKEN_EXPIRATION_DAYS")

	log.Println("config loaded for env:", env)

	return Config{
		Env: env,
		DatabaseDSN: databaseDSN,
		GoogleClientIDWeb: googleClientIDWeb,
		GoogleClientIDAndroid: googleClientIDAndroid,
		JWTExpirationMinutes: jwtExpirationMinutes,
		JWTRefreshExpirationMinutes: jwtRefreshExpirationDays * 24 * 60,
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s is not set", key)
	}
	return value
}

func mustGetInt(key string) int {
	value := mustGetEnv(key)

	number, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("%s must be a valid integer", key)
	}

	return number
}