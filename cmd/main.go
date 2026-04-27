package main

import (
	"go-tailwind-test/client"
	"go-tailwind-test/cmd/api"
	"go-tailwind-test/internal/config"
	"go-tailwind-test/internal/db"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/public", "public")

	cfg := config.Envs
	db, err := db.InitializePostgresDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	apiServer := api.NewAPIServer(db)
	if err := apiServer.APIService(app); err != nil {
		log.Fatal("API Server failed to initalize",err)
	}

	clientServer := client.NewClientServer(db)
	if err := clientServer.ClientService(app); err != nil {
		log.Fatal("Client Server failed to initalize",err)
	}

	if err := app.Start(":4200"); err != nil {
		log.Fatal(err)
	}
}

