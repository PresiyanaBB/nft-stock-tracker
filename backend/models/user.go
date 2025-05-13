package models

import (
	"github.com/google/uuid"
// 	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserRole string

const (
	Admin     UserRole = "admin"
	Collector UserRole = "collector"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email     string    `json:"email" gorm:"text;not null"`
	Role      UserRole  `json:"role" gorm:"text;default:collector"`
	Password  string    `json:"-" gorm:"not null"` // Do not compute the password in json
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BeforeCreate Automatically set a UUID before creating the record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.Email == "admin@gmail.com" {
		db.Model(u).Update("role", Admin)
	}

	return
}
