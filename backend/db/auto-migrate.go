package db

import (
	"github.com/PresiyanaBB/crypto-price-tracker/models"
	"gorm.io/gorm"
)

//func DBMigrator(db *gorm.DB) error {
//	// Drop the table if it exists
//	err := db.Migrator().DropTable(&models.NFT{})
//	if err != nil {
//		return err
//	}
//
//	// Recreate it with the updated struct
//	return db.AutoMigrate(&models.NFT{})
//}

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.NFT{})
}
