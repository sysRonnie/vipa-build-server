package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"github.com/joho/godotenv"
)

type Config struct {
	Env string 
	DatabaseDSN string 
}

var Envs = InitializeConfig()

func InitializeConfig() Config {
	env := strings.ToLower(os.Getenv("APP_ENV"))
	if env == "" {
		env = "dev"
	}

	loadDotEnv()

	databaseDSN := os.Getenv("DATABASE_DSN")
	log.Println("Database DSN loaded:", databaseDSN)

	return Config{
		Env: env, 
		DatabaseDSN: databaseDSN,
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