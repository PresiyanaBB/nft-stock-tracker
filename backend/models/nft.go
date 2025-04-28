package models

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type NFT struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	TokenID   uuid.UUID `json:"token_id"`
	TokenURI  string    `json:"token_uri"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	Price     float64   `json:"price"`
	Image     []byte    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NFTRepository interface {
	GetManyNFTs(ctx context.Context) ([]*NFT, error)
	GetNFT(ctx context.Context, nftId uuid.UUID) (*NFT, error)
	CreateNFT(ctx context.Context, nft NFT) (*NFT, error)
	UpdateNFT(ctx context.Context, nftId uuid.UUID) (*NFT, error)
	DeleteNFT(ctx context.Context, nftId uuid.UUID) error
}
