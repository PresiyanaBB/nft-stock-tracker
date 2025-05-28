package db

import (
	"github.com/PresiyanaBB/nft-stock-tracker/models"
	"github.com/PresiyanaBB/nft-stock-tracker/models/stock"
	"gorm.io/gorm"
)

// func DBMigrator(db *gorm.DB) error {
// 	// Drop the table if it exists
// 	err := db.Migrator().DropTable(&models.NFT{})
// 	if err != nil {
// 		return err
// 	}

// 	err_user := db.Migrator().DropTable(&models.User{})
// 	if err_user != nil {
// 		return err_user
// 	}

// 	err_usernft := db.Migrator().DropTable(&models.UserNFT{})
// 	if err_usernft != nil {
// 		return err_usernft
// 	}

// 	err_candle := db.Migrator().DropTable(&stock.Candle{})
// 	if err_candle != nil {
// 		return err_candle
// 	}

// 	// Recreate it with the updated struct
// 	return db.AutoMigrate(&models.NFT{}, models.UserNFT{}, models.User{}, &stock.Candle{})
// }

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.NFT{}, &models.UserNFT{}, &models.User{}, &stock.Candle{})
}
