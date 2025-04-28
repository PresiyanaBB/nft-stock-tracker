package repositories

import (
	"context"
	"github.com/PresiyanaBB/crypto-price-tracker/models"
	"github.com/google/uuid"
	"time"
)

type NFTRepository struct {
	db any
}

func (r *NFTRepository) GetManyNFTs(ctx context.Context) ([]*models.NFT, error) {
	nfts := []*models.NFT{}

	nfts = append(nfts, &models.NFT{
		ID:        uuid.New(),
		TokenID:   uuid.New(),
		TokenURI:  "hhtps://example.com/nft1",
		Name:      "NFT 1",
		Owner:     "0x1234567890abcdef",
		Price:     1000,
		Image:     []byte("image data"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

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

func NewNFTRepository(db any) models.NFTRepository {
	return &NFTRepository{
		db: db,
	}
}
