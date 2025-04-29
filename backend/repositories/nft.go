package repositories

import (
	"context"
	"github.com/PresiyanaBB/crypto-price-tracker/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NFTRepository struct {
	db *gorm.DB
}

func (r *NFTRepository) GetManyNFTs(ctx context.Context) ([]*models.NFT, error) {
	nfts := []*models.NFT{}

	res := r.db.Model(&models.NFT{}).Find(&nfts)

	if res.Error != nil {
		return nil, res.Error
	}

	return nfts, nil
}

func (r *NFTRepository) GetNFT(ctx context.Context, nftId uuid.UUID) (*models.NFT, error) {
	return nil, nil
}

func (r *NFTRepository) CreateNFT(ctx context.Context, nft models.NFT) (*models.NFT, error) {
	return nil, nil
}

func (r *NFTRepository) UpdateNFT(ctx context.Context, nftId uuid.UUID) (*models.NFT, error) {
	return nil, nil
}

func (r *NFTRepository) DeleteNFT(ctx context.Context, nftId uuid.UUID) error {
	return nil
}

func NewNFTRepository(db *gorm.DB) models.NFTRepository {
	return &NFTRepository{
		db: db,
	}
}
