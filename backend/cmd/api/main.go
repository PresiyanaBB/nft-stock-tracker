package main

import (
	"github.com/PresiyanaBB/crypto-price-tracker/handlers"
	"github.com/PresiyanaBB/crypto-price-tracker/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "CryptoPriceTracker",
		ServerHeader: "Fiber",
	})

	// Repositories
	nftRepository := repositories.NewNFTRepository(nil)

	//Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewNFTHandler(server.Group("/nft"), nftRepository)

	app.Listen(":3000")
}
