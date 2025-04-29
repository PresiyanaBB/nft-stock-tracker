package db

import (
	"github.com/PresiyanaBB/crypto-price-tracker/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.NFT{})
}
