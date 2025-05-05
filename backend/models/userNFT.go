package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserNFT struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	NFTID     uuid.UUID `json:"nft_id" gorm:"type:uuid;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	NFT       NFT       `json:"nft" gorm:"foreignkey:NFTID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Collected bool      `json:"collected" default:"false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserNFTRepository interface {
	GetManyUserNFTs(ctx context.Context, userId uuid.UUID) ([]*UserNFT, error)
	GetUserNFT(ctx context.Context, userId uuid.UUID, userNFTId uuid.UUID) (*UserNFT, error)
	CreateUserNFT(ctx context.Context, userId uuid.UUID, userNFT *UserNFT) (*UserNFT, error)
	UpdateUserNFT(ctx context.Context, userId uuid.UUID, userNFTId uuid.UUID, updateData map[string]interface{}) (*UserNFT, error)
}

type ValidateUserNFT struct {
	UserNFTId uuid.UUID `json:"UserNFTId"`
	OwnerId   uuid.UUID `json:"ownerId"`
}

// BeforeCreate Automatically set a UUID before creating the record
func (un *UserNFT) BeforeCreate(tx *gorm.DB) (err error) {
	if un.ID == uuid.Nil {
		un.ID = uuid.New()
	}
	return
}
