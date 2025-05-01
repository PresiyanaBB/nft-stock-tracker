package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type NFT struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey"`
	TokenURI  string          `json:"token_uri" gorm:"type:varchar(255)"`
	Name      string          `json:"name" gorm:"type:varchar(255)"`
	Creator   string          `json:"creator" gorm:"type:varchar(255)"`
	Price     decimal.Decimal `json:"price" gorm:"type:numeric(10,2)"`
	Image     []byte          `json:"image" gorm:"type:bytea"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}

// BeforeCreate Automatically set a UUID before creating the record
func (n *NFT) BeforeCreate(tx *gorm.DB) (err error) {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return
}

type NFTRepository interface {
	GetManyNFTs(ctx context.Context) ([]*NFT, error)
	GetNFT(ctx context.Context, nftId uuid.UUID) (*NFT, error)
	CreateNFT(ctx context.Context, nft *NFT) (*NFT, error)
	UpdateNFT(ctx context.Context, nftId uuid.UUID, updatedNFT map[string]interface{}) (*NFT, error)
	DeleteNFT(ctx context.Context, nftId uuid.UUID) error
}
