package main

import (
	"fmt"

	"github.com/PresiyanaBB/crypto-price-tracker/config"
	"github.com/PresiyanaBB/crypto-price-tracker/db"
	"github.com/PresiyanaBB/crypto-price-tracker/handlers"
	"github.com/PresiyanaBB/crypto-price-tracker/middlewares"
	"github.com/PresiyanaBB/crypto-price-tracker/repositories"
	"github.com/PresiyanaBB/crypto-price-tracker/services"
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
	userNFTRepository := repositories.NewUserNFTRepository(database)
	authRepository := repositories.NewAuthRepository(database)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(database))

	handlers.NewNFTHandler(privateRoutes.Group("/nft"), nftRepository)
	handlers.NewUserNFTHandler(privateRoutes.Group("/userNFT"), userNFTRepository)

	_ = app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
