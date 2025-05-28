package main

import (
	"fmt"

	"github.com/PresiyanaBB/nft-stock-tracker/config"
	"github.com/PresiyanaBB/nft-stock-tracker/db"
	"github.com/PresiyanaBB/nft-stock-tracker/handlers"
	"github.com/PresiyanaBB/nft-stock-tracker/handlers/stock"
	"github.com/PresiyanaBB/nft-stock-tracker/middlewares"
	"github.com/PresiyanaBB/nft-stock-tracker/repositories"
	"github.com/PresiyanaBB/nft-stock-tracker/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	envConfig := config.NewEnvConfig()
	database := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "CryptoPriceTracker",
		ServerHeader: "Fiber",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8082, http://192.168.43.150:8082, ws://localhost:8082, ws://192.168.43.150:8082, wss://ws.finnhub.io",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true, // Allow credentials (cookies, etc.)
	}))

	// Repositories
	nftRepository := repositories.NewNFTRepository(database)
	userNFTRepository := repositories.NewUserNFTRepository(database)
	authRepository := repositories.NewAuthRepository(database)
	candleRepository := repositories.NewCandleRepository(database)

	// Service
	authService := services.NewAuthService(authRepository)

	// WebSockets - Finnhub
	finnhubWSConn := stock.ConnectToFinnhub(envConfig)
	defer finnhubWSConn.Close()

	// Handle WebSocket
	go stock.HandleFinnhubMessages(finnhubWSConn, database)

	// Broadcast
	go stock.BroadcastUpdates()

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(database))

	handlers.NewCandleHandler(privateRoutes.Group("/candle"), candleRepository)
	handlers.NewNFTHandler(privateRoutes.Group("/nft"), nftRepository)
	handlers.NewUserNFTHandler(privateRoutes.Group("/userNFT"), userNFTRepository)
	app.Get("/ws", handlers.WSHandler)

	_ = app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
