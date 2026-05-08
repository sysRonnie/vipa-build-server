package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	env := strings.ToLower(os.Getenv("APP_ENV"))
	if env == "" {
		env = "dev"
	}

	loadDotEnv()

	databaseDSN := os.Getenv("DATABASE_DSN")
	googleClientIDWeb := os.Getenv("GOOGLE_CLIENT_ID_WEB")
	googleClientIDAndroid := os.Getenv("GOOGLE_CLIENT_ID_ANDROID")
	jwtExpirationMinutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_MINUTES"))
	if err != nil {
		log.Fatal(
			"invalid JWT_EXPIRATION_MINUTES",
		)
	}

	jwtRefreshExpirationDays, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_DAYS"))
	jwtRefreshExpirationMinutes := jwtRefreshExpirationDays * 24 * 60
	if err != nil {
		log.Fatal(
			"invalid REFRESH_TOKEN_EXPIRATION_DAYS",
		)
	}

	if googleClientIDWeb == "" || googleClientIDAndroid == "" {
		log.Fatal("Google Client IDs are not set in environment variables")
	}
	log.Println("Database DSN loaded:", databaseDSN)

	return Config{
		Env: env, 
		DatabaseDSN: databaseDSN,
		GoogleClientIDWeb: googleClientIDWeb,
		GoogleClientIDAndroid: googleClientIDAndroid,
		JWTExpirationMinutes: jwtExpirationMinutes,
		JWTRefreshExpirationMinutes: jwtRefreshExpirationMinutes,
	}
}

func loadDotEnv() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	for i := 0; i < 6; i++ {
		path := filepath.Join(dir, ".env")
		if _, err := os.Stat(path); err == nil {
			godotenv.Load(path)
			log.Println("✅ Loaded .env")
			return
		}
		dir = filepath.Dir(dir)
	}

	log.Println("⚠️  .env not found")
}