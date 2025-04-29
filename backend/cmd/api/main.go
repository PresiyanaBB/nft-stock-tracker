package main

import (
	"fmt"
	"github.com/PresiyanaBB/crypto-price-tracker/config"
	"github.com/PresiyanaBB/crypto-price-tracker/db"
	"github.com/PresiyanaBB/crypto-price-tracker/handlers"
	"github.com/PresiyanaBB/crypto-price-tracker/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	database := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "CryptoPriceTracker",
		ServerHeader: "Fiber",
	})

	// Repositories
	nftRepository := repositories.NewNFTRepository(database)

	//Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewNFTHandler(server.Group("/nft"), nftRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
