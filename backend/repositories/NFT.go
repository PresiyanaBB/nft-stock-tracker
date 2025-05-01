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

	res := r.db.Model(&models.NFT{}).Order("updated_at desc").Find(&nfts)

	if res.Error != nil {
		return nil, res.Error
	}

	return nfts, nil
}

func (r *NFTRepository) GetNFT(ctx context.Context, nftId uuid.UUID) (*models.NFT, error) {
	nft := &models.NFT{}

	res := r.db.Model(nft).Where("id = ?", nftId).First(nft)

	if res.Error != nil {
		return nil, res.Error
	}

	return nft, nil
}

func (r *NFTRepository) CreateNFT(ctx context.Context, nft *models.NFT) (*models.NFT, error) {
	res := r.db.Model(nft).Create(nft)

	if res.Error != nil {
		return nil, res.Error
	}

	return nft, nil
}

func (r *NFTRepository) UpdateNFT(ctx context.Context, nftId uuid.UUID, updatedNFT map[string]interface{}) (*models.NFT, error) {
	nft := &models.NFT{}

	updateRes := r.db.Model(nft).Where("id = ?", nftId).Updates(updatedNFT)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Model(nft).Where("id = ?", nftId).First(nft)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return nft, nil
}

func (r *NFTRepository) DeleteNFT(ctx context.Context, nftId uuid.UUID) error {
	res := r.db.Delete(&models.NFT{}, nftId)
	return res.Error
}

func NewNFTRepository(db *gorm.DB) *NFTRepository {
	return &NFTRepository{
		db: db,
	}
}
