package repositories

import (
	"context"

	"github.com/PresiyanaBB/crypto-price-tracker/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserNFTRepository struct {
	db *gorm.DB
}

func (r *UserNFTRepository) GetManyUserNFTs(ctx context.Context, userId uuid.UUID) ([]*models.UserNFT, error) {
	userNFTs := []*models.UserNFT{}

	res := r.db.Model(&models.UserNFT{}).Where("user_id = ?", userId).Preload("NFT").Order("updated_at desc").Find(&userNFTs)

	if res.Error != nil {
		return nil, res.Error
	}

	return userNFTs, nil
}

func (r *UserNFTRepository) GetAllUserNFTs(ctx context.Context) ([]*models.UserNFT, error) {
	userNFTs := []*models.UserNFT{}

	res := r.db.Model(&models.UserNFT{}).Preload("NFT").Order("updated_at desc").Find(&userNFTs)

	if res.Error != nil {
		return nil, res.Error
	}

	return userNFTs, nil
}

func (r *UserNFTRepository) GetUserNFT(ctx context.Context, userId uuid.UUID, userNFTId uuid.UUID) (*models.UserNFT, error) {
	userNFT := &models.UserNFT{}

	res := r.db.Model(&models.UserNFT{}).Where("id = ? and user_id = ?", userNFTId, userId).Preload("NFT").First(userNFT)

	if res.Error != nil {
		return nil, res.Error
	}

	return userNFT, nil
}

func (r *UserNFTRepository) CreateUserNFT(ctx context.Context, userId uuid.UUID, userNFT *models.UserNFT) (*models.UserNFT, error) {
	userNFT.UserID = userId

	var existingUserNFT models.UserNFT
	if err := r.db.Model(&models.UserNFT{}).
		Where("user_id = ? AND nft_id = ?", userId, userNFT.NFTID).
		First(&existingUserNFT).Error; err == nil {
		return nil, err
	}

	res := r.db.Model(userNFT).Create(userNFT)

	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetUserNFT(ctx, userId, userNFT.ID)
}

func (r *UserNFTRepository) UpdateUserNFT(ctx context.Context, userId uuid.UUID, userNFTId uuid.UUID, updateData map[string]interface{}) (*models.UserNFT, error) {
	userNFT := &models.UserNFT{}

	updateRes := r.db.Model(userNFT).Where("id = ?", userNFTId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	return r.GetUserNFT(ctx, userId, userNFTId)
}

func (r *UserNFTRepository) DeleteUserNFT(ctx context.Context, userNFTId uuid.UUID) error {
	userNFT := &models.UserNFT{}

	res := r.db.Model(userNFT).Where("id = ?", userNFTId).Delete(userNFT)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func NewUserNFTRepository(db *gorm.DB) *UserNFTRepository {
	return &UserNFTRepository{
		db: db,
	}
}
