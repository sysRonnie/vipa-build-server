package main

import (
	"go-tailwind-test/internal/config"
	"go-tailwind-test/internal/db"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Initialize PostgreSQL database connection
	db, err := db.InitializePostgresDB(config.Envs.DatabaseDSN)
	//db, err := DB.NewPostgresStorage(config.Envs.LocalPostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	// Configure the PostgreSQL driver for migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // Path to your migrations folder
		"postgres",                     // Database name (matches the driver)
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Check the command-line argument to determine direction (up/down/force)
	if len(os.Args) < 2 {
		log.Println("Usage: go run main.go [up|down|force <version>]")
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations applied successfully.")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations rolled back successfully.")
	case "force":
		if len(os.Args) < 3 {
			log.Println("Usage: go run main.go force <version>")
			return
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid version: %v", err)
		}
		if err := m.Force(version); err != nil {
			log.Fatal(err)
		}
		log.Printf("Forced migration to version %d successfully.\n", version)
	default:
		log.Println("Usage: go run main.go [up|down|force <version>]")
	}
}