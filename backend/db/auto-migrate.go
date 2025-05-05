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
//	err_user := db.Migrator().DropTable(&models.User{})
//	if err_user != nil {
//		return err_user
//	}
//
//	err_usernft := db.Migrator().DropTable(&models.UserNFT{})
//	if err_usernft != nil {
//		return err_usernft
//	}
//
//	// Recreate it with the updated struct
//	return db.AutoMigrate(&models.NFT{}, models.UserNFT{}, models.User{})
//}

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.NFT{}, &models.UserNFT{}, &models.User{})
}
